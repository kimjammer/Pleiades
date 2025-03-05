package main

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Command interface {
	apply(project *Project)
}

type ConnectionForSpace struct {
	state_tx   chan<- []byte
	command_rx <-chan Command
}

type ConnectionForSocket struct {
	state_rx   <-chan []byte
	command_tx chan<- Command
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
		if contactPoint, exists := projectSpaces[project]; exists {
			contactPoint.requeryRequest <- struct{}{}
		}
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

	state_chan := make(chan []byte)
	command_chan := make(chan Command)

	projectSpaces[projectId].sendNewConnection <- ConnectionForSpace{state_tx: state_chan, command_rx: command_chan}

	return ConnectionForSocket{
		state_rx:   state_chan,
		command_tx: command_chan,
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

	var connections []ConnectionForSpace

	for {
		select {
		case newConnection := <-connectionsChannel:
			newConnection.state_tx <- encodeProject(project, users)
			connections = append(connections, newConnection)
		case _ = <-noConnectionsTimeout:
			if len(connections) == 0 {
				return
			}
		case _ = <-requeryChannel:
			users = queryUsers(project.Users)
		default:
		}

		appliedCommand := false

		i := len(connections) - 1
		for i >= 0 {
			connection := connections[i]

			select {
			case command, ok := <-connection.command_rx:
				if !ok {
					// Channel closed, stop listening to it
					connections[i] = connections[len(connections)-1]
					connections = connections[:len(connections)-1]
					i -= 1

					if len(connections) == 0 {
						go timeout(noConnectionsTimeout)
					}

					continue
				}

				command.apply(&project)

				if _, is := command.(UserLeave); is {
					users = queryUsers(project.Users)
					if !existingInviteLinks(project) && maybeDeleteProject(project) {
						return
					}
				}

				if _, is := command.(Delete); is {
					users = queryUsers(project.Users)
					if maybeDeleteProject(project) {
						return
					}
				}

				appliedCommand = true
			default:
			}

			i -= 1
		}

		if appliedCommand {
			encoded := encodeProject(project, users)

			for _, connection := range connections {
				connection.state_tx <- encoded
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
	// TODO
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
