package main

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// I don't know how to test with a real websocket connection so this will have to do. It will be pretty realistic anyways. The playwright tests will probably implicitly test it too.
type TestConn struct {
	tx         chan<- string
	rx         <-chan string
	closeCheck chan struct{}
}

func (self TestConn) isClosed() bool {
	select {
	case _, _ = <-self.closeCheck:
		return true
	default:
		return false
	}
}

func (self TestConn) send(data string) bool {
	self.tx <- data
	return self.isClosed()
}

func (self TestConn) recv() (string, bool) {
	data := <-self.rx

	return data, self.isClosed()
}

func (self TestConn) close() {
	close(self.tx)
	close(self.closeCheck)
	_, alive := <-self.rx
	if alive {
		panic("The server should close the connection upon the client closing")
	}
}

func (self TestConn) recvState() (ProjectAndUsers, bool) {
	var project ProjectAndUsers

	data, closed := self.recv()
	if closed {
		return project, true
	}

	if data == "PROJECT ID DNE" {
		panic("Project ID DNE")
	}

	if strings.HasPrefix(data, "FAIL:") {
		panic(data)
	}

	err := json.Unmarshal(([]byte)(data), &project)
	if err != nil {
		panic(err)
	}

	return project, false
}

func connect(t *testing.T, userId string, projectId string) (TestConn, ProjectAndUsers) {
	a := make(chan string)
	b := make(chan string)
	closeCheck := make(chan struct{})

	forConnHandler := TestConn{tx: a, rx: b, closeCheck: closeCheck}
	forReturn := TestConn{tx: b, rx: a, closeCheck: closeCheck}

	go handleConnection(forConnHandler, userId)

	forReturn.send(projectId)

	state, disconnect := forReturn.recvState()
	require.False(t, disconnect)

	return forReturn, state
}

// USER STORY 2.6

func TestLeavingDeletesProject(t *testing.T) {
	resetDB()

	conn, _ := connect(t, "67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

	conn.send(`{
		"Name": "leave",
		"Args": {}
	}`)

	_, disconnect := conn.recvState()
	require.True(t, disconnect)

	user, err := getUserById(context.TODO(), "67c7b20021675682e4395270")
	require.Nil(t, err)
	require.Equal(t, len(user.Projects), 0)

	// With no invite links around, leaving a project as the only user should delete it
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)
}
