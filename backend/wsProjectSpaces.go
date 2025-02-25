package main

import (
	"context"
	"encoding/json"
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

var projectSpacesMutex sync.Mutex
var projectSpaces = make(map[string]chan<- ConnectionForSpace)

func joinSpace(projectId string) ConnectionForSocket {
	projectSpacesMutex.Lock()
	defer projectSpacesMutex.Unlock()

	if _, ok := projectSpaces[projectId]; !ok {
		channel := make(chan ConnectionForSpace)
		go projectSpace(channel, projectId)
		projectSpaces[projectId] = channel
	}

	state_chan := make(chan []byte)
	command_chan := make(chan Command)

	projectSpaces[projectId] <- ConnectionForSpace{state_tx: state_chan, command_rx: command_chan}

	return ConnectionForSocket{
		state_rx:   state_chan,
		command_tx: command_chan,
	}
}

func encodeProject(project Project) []byte {
	encoded, err := json.Marshal(project)
	if err != nil {
		panic(err)
	}
	return encoded
}

func projectSpace(connectionsChannel <-chan ConnectionForSpace, projectId string) {
	defer func() {
		projectSpacesMutex.Lock()
		defer projectSpacesMutex.Unlock()

		delete(projectSpaces, projectId)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: projectId}}
	var project Project
	err := db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		panic(err)
	}

	var connections []ConnectionForSpace

	for {
		select {
		case newConnection := <-connectionsChannel:
			newConnection.state_tx <- encodeProject(project)
			connections = append(connections, newConnection)
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
					continue
				}

				command.apply(&project)
				appliedCommand = true
			}

			i -= 1
		}

		if appliedCommand {
			encoded := encodeProject(project)

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
