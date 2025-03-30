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

	go handleConnection(WsConn{socket: conn}, userId.(string))
}

type Connection interface {
	send(data string) bool
	recv() (string, bool)
	close()
}

type WsConn struct {
	socket *websocket.Conn
}

func (self WsConn) send(data string) bool {
	if err := self.socket.WriteMessage(websocket.TextMessage, []byte(data)); err != nil {
		if _, ok := err.(*websocket.CloseError); ok {
			return true
		}
		panic(err)
	}

	return false
}

func (self WsConn) recv() (string, bool) {
	mt, bytes, err := self.socket.ReadMessage()
	if err != nil {
		if _, ok := err.(*websocket.CloseError); ok {
			return "", true
		}
		panic(err)
	}

	str := string(bytes)

	if mt != websocket.TextMessage {
		if err := self.socket.WriteMessage(websocket.TextMessage, []byte("FAIL: All websockets messages must be text")); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return "", true
			}
			panic(err)
		}
	}

	return str, false
}

func (self WsConn) close() {
	self.socket.Close()
}

func handleConnection(conn Connection, userId string) {
	disconnectSignal := make(chan struct{})
	defer func() {
		// Don't close until the wsProjectSpace has closed
		<-disconnectSignal
		conn.close()
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	userObjectId, _ := primitive.ObjectIDFromHex(userId)
	filter := bson.D{{Key: "_id", Value: userObjectId}}
	var crrUser User
	err := db.Collection("users").FindOne(ctx, filter).Decode(&crrUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			disconnect := conn.send("FAIL: User ID not found")
			if disconnect {
				return
			}
		} else {
			panic(err)
		}
	}

	projectId, disconnect := conn.recv()
	if disconnect {
		return
	}

	if !slices.Contains(crrUser.Projects, projectId) {
		conn.send("PROJECT ID DNE")
	}

	projectSpace := joinSpace(projectId)
	defer func() {
		close(projectSpace.command_tx)
	}()

	go func() {
		defer func() { disconnectSignal <- struct{}{} }()

		for {
			select {
			case newState, ok := <-projectSpace.state_rx:
				if !ok {
					projectSpace.state_rx = nil
					continue
				}
				disconnect := conn.send(string(newState))
				if disconnect {
					return
				}

			case err, ok := <-projectSpace.errors:
				{
					if !ok {
						return
					}
					if err == nil {
						continue
					}

					log.Println(err.Error())

					disconnect := conn.send("FAIL: " + err.Error())
					if disconnect {
						return
					}
				}
			}
		}
	}()

	//Listen loop
	for {
		message, disconnect := conn.recv()
		if disconnect {
			return
		}

		command := CommandMessage{}

		err = json.Unmarshal(([]byte)(message), &command)
		if err != nil {
			projectSpace.errors <- err
			continue
		}

		decodedCommand, err := decodeCommand(command, userId)

		if err != nil {
			projectSpace.errors <- err
			continue
		}

		if command.Name == "leave" || command.Name == "delete" {
			leave(userId, projectId)

			projectSpace.command_tx <- decodedCommand

			return
		}

		projectSpace.command_tx <- decodedCommand
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
