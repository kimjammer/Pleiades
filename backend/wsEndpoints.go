package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*
TODO: Completely untested - John
*/

func wsEndpoint(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	mt, key, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message: ", err)
		return
	}

	if mt != websocket.TextMessage {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: Expected the key to be a text message")); err != nil {
			log.Println("Error writing message: ", err)
			return
		}
	}

	log.Println(key)

	mt, projectId, err := conn.ReadMessage()
	if err != nil {
		log.Println("Error reading message: ", err)
		return
	}

	if mt != websocket.TextMessage {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("FAIL: Expected the project ID to be a text message")); err != nil {
			log.Println("Error writing message: ", err)
			return
		}
	}

	log.Println(key)
	log.Println(projectId)

	//Listen loop
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			return
		}

		//DEBUG
		log.Println("Received: ", string(message))

		//DEBUG: Echo same message back
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("Error writing message: ", err)
			break
		}
	}
}
