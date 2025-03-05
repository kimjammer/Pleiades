package main

import (
	"context"
	"encoding/json"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:5173"
	},
}

func wsEndpoint(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		conn.Close()
		log.Println(err)
		return
	}

	userId, exists := c.Get("userId")

	if !exists {
		defer conn.Close()

		if err := conn.WriteMessage(websocket.TextMessage, []byte("UNAUTHORIZED")); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			panic(err)
		}

		return
	}

	go handleConnection(conn, userId.(string))
}

func handleConnection(conn *websocket.Conn, userId string) {
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userObjectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: userObjectId}}
	var crrUser User
	err := db.Collection("users").FindOne(ctx, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: User ID not found")); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	mt, projectIdBytes, err := conn.ReadMessage()
	if err != nil {
		if _, ok := err.(*websocket.CloseError); ok {
			return
		}
		panic(err)
	}

	projectId := string(projectIdBytes)

	if mt != websocket.TextMessage {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: Expected the project ID to be a text message")); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			panic(err)
		}
	}

	if !slices.Contains(crrUser.Projects, projectId) {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("PROJECT ID DNE")); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			panic(err)
		}
	}

	projectSpace := joinSpace(projectId)
	disconnectSignal := make(chan struct{})
	defer func() {
		disconnectSignal <- struct{}{}
		close(projectSpace.command_tx)
	}()

	go func() {
		for {
			select {
			case newState := <-projectSpace.state_rx:
				if err := conn.WriteMessage(websocket.TextMessage, newState); err != nil {
					if _, ok := err.(*websocket.CloseError); ok {
						return
					}
					panic(err)
				}

			case _ = <-disconnectSignal:
				return
			}
		}
	}()

	//Listen loop
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			panic(err)
		}

		if mt != websocket.TextMessage {
			if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: Expected the command to be a text message")); err != nil {
				if _, ok := err.(*websocket.CloseError); ok {
					return
				}
				panic(err)
			}
		}

		command := CommandMessage{}

		err = json.Unmarshal(message, &command)
		if err != nil {
			panic(err)
		}

		decodedCommand := decodeCommand(command, userId)

		projectSpace.command_tx <- decodedCommand

		if command.Name == "leave" || command.Name == "delete" {
			// TODO: Delete if there are no other members and no active invite links

			leave(userId, projectId)

			return
		}
	}
}

func leave(userId string, projectId string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userObjectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: userObjectId}}
	update := bson.D{{"$pull", bson.D{{"projects", projectId}}}}
	_, err := db.Collection("users").UpdateOne(ctx, filter, update)
	if err != nil {
		panic(err)
	}
}

type CommandMessage struct {
	Name string
	Args json.RawMessage
}

type DemoButtonCommand struct {
	newState string
}

func (self DemoButtonCommand) apply(state *Project) {
	if self.newState == "" {
		state.DemoButtonState = ""
	} else if state.DemoButtonState == "" && slices.Contains([]string{"a", "b"}, self.newState) {
		state.DemoButtonState = self.newState
	}
}

type UserLeave struct {
	userId string
}

func (self UserLeave) apply(state *Project) {
	for i, user := range state.Users {
		if user.User == self.userId {
			state.Users[i].LeftProject = true
		}
	}
}

type Delete struct {
	userId string
}

func (self Delete) apply(state *Project) {
	UserLeave(self).apply(state)
}

func decodeCommand(command CommandMessage, userId string) Command {
	if command.Name == "demoButtonState" {
		var newState string
		err := json.Unmarshal(command.Args, &newState)
		if err != nil {
			panic("Client sent invalid command")
		}

		return DemoButtonCommand{newState}
	}

	if command.Name == "leave" {
		return UserLeave{userId: userId}
	}

	if command.Name == "delete" {
		return Delete{userId: userId}
	}

	log.Println("Command:", command)
	panic("Unknown command:" + command.Name)
}
