package main

import (
	"context"
	"encoding/json"
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
	Id           string
	LeftProject  bool
	FirstName    string
	LastName     string
	Availability []Availability
}

var projectSpacesMutex sync.Mutex
var projectSpaces = make(map[string]ProjectSpaceContactPoint)

// When a change is applied to a user (availability), this function will retransmit the data to all of the projects that the user is a member of
func requeryUser(user User) {
	projectSpacesMutex.Lock()
	defer projectSpacesMutex.Unlock()

	for _, project := range user.Projects {
		requeryUsersForProject(project)
	}
}

// Requery the users for a particular project ID
func requeryUsersForProject(projectId string) {
	if contactPoint, exists := projectSpaces[projectId]; exists {
		contactPoint.requeryRequest <- struct{}{}
	}
}

func joinSpace(projectId string) ConnectionForSocket {
	projectSpacesMutex.Lock()
	defer projectSpacesMutex.Unlock()

	if _, ok := projectSpaces[projectId]; !ok {
		contactPoint := ProjectSpaceContactPoint{
			sendNewConnection: make(chan ConnectionForSpace),
			requeryRequest:    make(chan struct{}),
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
	Project Project
	Users   []UserInProject
}

func encodeProject(project Project, users []UserInProject) []byte {
	encoded, err := json.Marshal(ProjectAndUsers{
		Project: project,
		Users:   users,
	})
	if err != nil {
		panic(err)
	}
	return encoded
}

func timeout(channel chan<- struct{}) {
	secs := 10.

	duration := time.Duration(secs * 1000 * 1000 * 1000)

	if duration.Seconds() != secs {
		panic("Conversion wrong")
	}

	time.Sleep(duration)

	channel <- struct{}{}
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

	var noConnectionsTimeout = make(chan struct{})

	go timeout(noConnectionsTimeout)

	projectStateFanOut := []StateFanOutTx{}
	var commandChannel chan FanInMessage = make(chan FanInMessage)
	connectionCount := 0

	for {
		appliedChange := false

		select {
		case newConnection := <-connectionsChannel:
			connectionCount += 1
			// Guaranteed not to block since it's the first one and there's a buffer of one
			newConnection.state_tx <- encodeProject(project, users)
			stateChannel := make(chan []byte, 1)
			killedChannel := make(chan interface{})
			go fanInCommandsFrom(newConnection, commandChannel, StateFanOutRecv{stateChannel, killedChannel})
			projectStateFanOut = append(projectStateFanOut, StateFanOutTx{stateChannel, killedChannel})
		case _ = <-noConnectionsTimeout:
			if connectionCount == 0 {
				return
			}
		case _ = <-requeryChannel:
			users = queryUsers(project.Users)
			appliedChange = true
		case message := <-commandChannel:
			if message.wasKilled {
				connectionCount -= 1

				if connectionCount == 0 {
					go timeout(noConnectionsTimeout)
				}

				continue
			}

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
						return false, true
					}
				}

				if _, is := message.command.(Delete); is {
					users = queryUsers(project.Users)
					if maybeDeleteProject(project) {
						return false, true
					}
				}

				return true, false
			})()

			if deleted {
				return
			}
		}

		if appliedChange {
			encoded := encodeProject(project, users)

			projectStateFanOut = slices.DeleteFunc(projectStateFanOut, func(v StateFanOutTx) bool {
				select {
				case _ = <-v.killed:
					return true

				default:
					return false
				}
			})

			for _, fanOut := range projectStateFanOut {
				fanOut.recvStates <- encoded
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
	recvStates chan<- []byte
	killed     <-chan interface{}
}

func fanInCommandsFrom(connection ConnectionForSpace, sendTo chan<- FanInMessage, recvStates StateFanOutRecv) {
	defer func() {
		close(recvStates.killed)
	}()

	for {
		select {
		case command, ok := <-connection.command_rx:
			if !ok {
				sendTo <- FanInMessage{command: nil, connection: connection, wasKilled: true}
				return
			}

			sendTo <- FanInMessage{command: command, connection: connection, wasKilled: false}
		case recvState := <-recvStates.recvStates:
			select {
			case connection.state_tx <- recvState:
			default:
				// Flush the channel to remove the un-received state
				<-connection.state_tx
				connection.state_tx <- recvState
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

		userInProject := UserInProject{
			Id:           user.User,
			LeftProject:  user.LeftProject,
			FirstName:    userInDB.FirstName,
			LastName:     userInDB.LastName,
			Availability: userInDB.Availability,
		}

		out = append(out, userInProject)
	}

	return out
}

func existingInviteLinks(project Project) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	thing := db.Collection("invitations").FindOne(ctx, bson.D{{Key: "ProjectId", Value: project.Id}})

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
