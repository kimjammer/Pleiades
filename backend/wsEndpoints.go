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
		allowedOrigins := []string{"http://localhost:5173", "http://localhost:4173"}
		return slices.Contains(allowedOrigins, r.Header.Get("Origin"))
	},
}

func wsEndpoint(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		conn.Close()
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

			case err := <-projectSpace.errors:
				{
					log.Println(err.Error())

					if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: "+err.Error())); err != nil {
						if _, ok := err.(*websocket.CloseError); ok {
							return
						}
						panic(err)
					}
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

			projectSpace.errors <- err
		}

		if mt != websocket.TextMessage {
			projectSpace.errors <- errors.New("Expected the command to be a text message")
			continue
		}

		command := CommandMessage{}

		err = json.Unmarshal(message, &command)
		if err != nil {
			projectSpace.errors <- err
			continue
		}

		decodedCommand, err := decodeCommand(command, userId)

		if err != nil {
			projectSpace.errors <- err
			continue
		}

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
