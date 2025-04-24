package main

import (
	"context"
	"encoding/json"
	"errors"
	"slices"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ConnectionForSpace struct {
	// This channel must be "two way" because the project space goroutine needs to clear the state if the buffer is full
	state_tx   chan []byte
	command_rx <-chan Command
	error_tx   chan<- error
}

type ConnectionForSocket struct {
	state_rx   <-chan []byte
	command_tx chan<- Command
	errors     chan error // This channel is two-way because it's convenient for the socket to send itself errors
}

type ProjectSpaceContactPoint struct {
	sendNewConnection chan ConnectionForSpace
	requeryRequest    chan struct{}
}

type UserInProject struct {
	Id            string
	LeftProject   bool
	FirstName     string
	LastName      string
	Email         string
	Availability  []Availability
	NotifSettings []bool
}

var projectSpacesMutex sync.Mutex
var projectSpaces = make(map[string]ProjectSpaceContactPoint)

// When a change is applied to a user (availability), this function will retransmit the data to all of the projects that the user is a member of
func requeryUser(user User) {
	for _, project := range user.Projects {
		requeryUsersForProject(project)
	}
}

// Requery the users for a particular project ID
func requeryUsersForProject(projectId string) {
	projectSpacesMutex.Lock()
	defer projectSpacesMutex.Unlock()

	if contactPoint, exists := projectSpaces[projectId]; exists {
		contactPoint.requeryRequest <- struct{}{}
	}
}

func joinSpace(projectId string) ConnectionForSocket {
	projectSpacesMutex.Lock()
	defer projectSpacesMutex.Unlock()

	if _, ok := projectSpaces[projectId]; !ok {
		contactPoint := ProjectSpaceContactPoint{
			sendNewConnection: make(chan ConnectionForSpace, 1),
			requeryRequest:    make(chan struct{}, 1),
		}
		go projectSpace(contactPoint, projectId)
		projectSpaces[projectId] = contactPoint
	}

	state_chan := make(chan []byte, 1)
	command_chan := make(chan Command, 1)
	error_chan := make(chan error, 1)

	projectSpaces[projectId].sendNewConnection <- ConnectionForSpace{state_tx: state_chan, command_rx: command_chan, error_tx: error_chan}

	return ConnectionForSocket{
		state_rx:   state_chan,
		command_tx: command_chan,
		errors:     error_chan,
	}
}

type ProjectAndUsers struct {
	Project      Project
	Users        []UserInProject
	Notification *Notify
}

func encodeProject(project Project, users []UserInProject, notif *Notify) []byte {
	encoded, err := json.Marshal(ProjectAndUsers{
		Project:      project,
		Users:        users,
		Notification: notif,
	})
	if err != nil {
		panic(err)
	}
	return encoded
}

func timeout(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	secs := 10.

	duration := time.Duration(secs * 1000 * 1000 * 1000)

	if duration.Seconds() != secs {
		panic("Conversion wrong")
	}

	time.Sleep(duration)
}

func projectSpace(contactPoint ProjectSpaceContactPoint, projectId string) {
	defer func() {
		projectSpacesMutex.Lock()
		defer projectSpacesMutex.Unlock()

		delete(projectSpaces, projectId)
	}()

	connectionsChannel := contactPoint.sendNewConnection
	requeryChannel := contactPoint.requeryRequest

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: projectId}}
	var project Project
	err := db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		panic(err)
	}

	users := queryUsers(project.Users)

	var wg sync.WaitGroup

	go timeout(&wg)

	projectStateFanOut := []StateFanOutTx{}
	var commandChannel chan FanInMessage = make(chan FanInMessage, 16)

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	defer func() {
		for _, conn := range projectStateFanOut {
			close(conn.txStates)
		}
	}()

	for {
		appliedChange := false
		var notif *Notify = nil

		select {
		case newConnection := <-connectionsChannel:
			// Guaranteed not to block since it's the first one and there's a buffer of one
			newConnection.state_tx <- encodeProject(project, users, nil)
			stateChannel := make(chan []byte, 1)
			killedChannel := make(chan interface{})
			go fanInCommandsFrom(newConnection, commandChannel, StateFanOutRecv{stateChannel, killedChannel}, &wg)
			projectStateFanOut = append(projectStateFanOut, StateFanOutTx{stateChannel, killedChannel})
		case _ = <-done:
			return
		case _ = <-requeryChannel:
			users = queryUsers(project.Users)
			appliedChange = true
		case message := <-commandChannel:
			if notify, ok := message.command.(Notify); ok {
				notif = &notify
				appliedChange = true
			} else {
				maybe_err := message.command.apply(&project)

				var deleted bool
				appliedChange, deleted = (func() (bool, bool) {
					defer (func() {
						// Send the error even if it's nil, to allow `applyCommandToProject` to possibly return an error
						// Also this needs to be done after project deletion so that tests see an updated database
						message.connection.error_tx <- maybe_err
					})()

					if maybe_err != nil {
						return false, false
					}

					if _, is := message.command.(UserLeave); is {
						users = queryUsers(project.Users)
						hasInviteLinks := existingInviteLinks(project)
						if !hasInviteLinks && maybeDeleteProject(project) {
							return true, true
						}
					}

					if _, is := message.command.(Delete); is {
						users = queryUsers(project.Users)
						if maybeDeleteProject(project) {
							return true, true
						} else {
							maybe_err = errors.New("Cannot delete when there are other users in the project")
						}
					}

					return true, false
				})()

				if deleted {
					return
				}
			}
		}

		if appliedChange {
			encoded := encodeProject(project, users, notif)

			projectStateFanOut = slices.DeleteFunc(projectStateFanOut, func(v StateFanOutTx) bool {
				select {
				case _ = <-v.killed:
					close(v.txStates)
					return true

				default:
					return false
				}
			})

			for _, fanOut := range projectStateFanOut {
				fanOut.txStates <- encoded
			}

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			update := bson.D{{"$set", project}}
			_, err := db.Collection("projects").UpdateOne(ctx, filter, update)
			if err != nil {
				panic(err)
			}
		}
	}
}

type FanInMessage struct {
	command    Command
	connection ConnectionForSpace
	wasKilled  bool
}

type StateFanOutRecv struct {
	recvStates <-chan []byte
	killed     chan<- interface{}
}

type StateFanOutTx struct {
	txStates chan<- []byte
	killed   <-chan interface{}
}

func fanInCommandsFrom(connection ConnectionForSpace, sendTo chan<- FanInMessage, recvStates StateFanOutRecv, wg *sync.WaitGroup) {
	wg.Add(1)
	defer func() {
		close(connection.state_tx)
		close(connection.error_tx)
		close(recvStates.killed)
		wg.Done()
	}()

	defunct := false

	for {
		select {
		case command, ok := <-connection.command_rx:
			if !ok {
				connection.command_rx = nil
				defunct = true
				continue
			}

			sendTo <- FanInMessage{command: command, connection: connection, wasKilled: false}
		case recvState, ok := <-recvStates.recvStates:
			if !ok {
				return
			}

			select {
			case connection.state_tx <- recvState:
			default:
				// Flush the channel to remove the un-received state
				<-connection.state_tx
				connection.state_tx <- recvState
			}

			if defunct {
				return
			}
		}
	}
}

func queryUsers(users []UserAndLeft) []UserInProject {
	out := []UserInProject{}

	for _, user := range users {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		userObjectId, _ := primitive.ObjectIDFromHex(user.User)
		filter := bson.D{{Key: "_id", Value: userObjectId}}
		var userInDB User
		err := db.Collection("users").FindOne(ctx, filter).Decode(&userInDB)
		if err != nil {
			panic(err)
		}

		if userInDB.NotifSettings == nil {
			userInDB.NotifSettings = []bool{false, false, false}
		}

		userInProject := UserInProject{
			Id:            user.User,
			LeftProject:   user.LeftProject,
			FirstName:     userInDB.FirstName,
			LastName:      userInDB.LastName,
			Email:         userInDB.Email,
			Availability:  userInDB.Availability,
			NotifSettings: userInDB.NotifSettings,
		}

		out = append(out, userInProject)
	}

	return out
}

func existingInviteLinks(project Project) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	thing := db.Collection("invitations").FindOne(ctx, bson.D{{Key: "projectid", Value: project.Id}})

	if thing.Err() == mongo.ErrNoDocuments {
		return false
	}

	if thing.Err() != nil {
		panic(thing.Err())
	}

	return true
}

func maybeDeleteProject(project Project) bool {
	for _, user := range project.Users {
		if !user.LeftProject {
			return false
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: project.Id}}
	_, err := db.Collection("projects").DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}

	return true
}
