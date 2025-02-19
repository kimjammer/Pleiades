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
		return
	}
	defer conn.Close()

	//Listen loop
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			break
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
