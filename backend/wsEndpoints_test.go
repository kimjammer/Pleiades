package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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
	data, ok := <-self.rx

	return data, !ok
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

func TestLeavingInviteAroundAndUserAround(t *testing.T) {
	resetDB()

	router := setupTestRouter()

	// Create an invite link
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/invite?id=53ed4d28-9279-4b4e-9256-b1e693332625", nil)
	router.ServeHTTP(w, req)
	bodyBytes, _ := io.ReadAll(w.Body)
	body := string(bodyBytes)
	require.Equal(t, http.StatusOK, w.Code)

	conn, _ := connect(t, "67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

	conn.send(`{
		"Name": "leave",
		"Args": {}
	}`)

	project, disconnect := conn.recvState()
	require.False(t, disconnect)

	require.Equal(t, len(project.Users), 1)
	require.True(t, project.Users[0].LeftProject)

	_, disconnect = conn.recvState()
	require.True(t, disconnect)

	user, err := getUserById(context.TODO(), "67c7b20021675682e4395270")
	require.Nil(t, err)
	require.Equal(t, len(user.Projects), 0)

	// There is still an invite link, so don't delete it
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)

	// Rejoin the project
	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/join?id="+body, nil)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	conn, project = connect(t, "67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

	require.Equal(t, len(project.Users), 1)
	require.False(t, project.Users[0].LeftProject)

	// Make another person join

	secondPerson()

	routerTwo := setupTestRouterCustomToken("67e5b7abef0709e1426eed50")

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/join?id="+body, nil)
	routerTwo.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	// First update from joining (updateProject):
	_, _ = conn.recvState()

	// Second update from joining (requeryUsers):
	project, disconnect = conn.recvState()
	require.False(t, disconnect)
	require.Equal(t, len(project.Users), 2)
	require.False(t, project.Users[0].LeftProject)
	require.False(t, project.Users[1].LeftProject)

	// Get rid of the invite to test that having another user around doesn't delete
	_, err = db.Collection("invitations").DeleteOne(context.TODO(), bson.D{{Key: "projectid", Value: "53ed4d28-9279-4b4e-9256-b1e693332625"}})
	require.Nil(t, err)

	count, _ = db.Collection("invitations").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)

	// Make them leave the project
	conn2, project2 := connect(t, "67e5b7abef0709e1426eed50", "53ed4d28-9279-4b4e-9256-b1e693332625")

	require.Equal(t, len(project2.Users), 2)

	conn2.send(`{
		"Name": "leave",
		"Args": {}
	}`)

	project2, disconnect = conn2.recvState()
	require.False(t, disconnect)
	require.Equal(t, len(project2.Users), 2)
	require.False(t, project2.Users[0].LeftProject)
	require.True(t, project2.Users[1].LeftProject)

	_, disconnect = conn2.recvState()
	require.True(t, disconnect)

	// Update from the person leaving:
	project, disconnect = conn.recvState()
	require.False(t, disconnect)
	require.Equal(t, len(project.Users), 2)
	require.False(t, project.Users[0].LeftProject)
	require.True(t, project.Users[1].LeftProject)

	// There is another user in the project, so don't delete it
	count, _ = db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)

	// Make the first user leave
	conn.send(`{
		"Name": "leave",
		"Args": {}
	}`)

	project, disconnect = conn.recvState()

	_, disconnect = conn.recvState()
	require.True(t, disconnect)

	// Now there are no users, so delete the project
	count, _ = db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)
}
