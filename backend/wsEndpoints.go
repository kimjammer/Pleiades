package main

import (
	"context"
	"errors"
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
		log.Println(err)
		return
	}

	go handleConnection(conn)
}

func handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	mt, token_bytes, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}
	token := string(token_bytes)

	if mt != websocket.TextMessage {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: Expected the token to be a text message")); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			panic(err)
		}
	}

	userId := verify_token(&token)
	if userId == nil {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("INVALID TOKEN")); err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				return
			}
			panic(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{Key: "_id", Value: userId}}
	var crrUser User
	err = db.Collection("users").FindOne(ctx, filter).Decode(&crrUser)
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

		log.Println(message)
	}
}
