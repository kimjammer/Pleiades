package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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

func (self TestConn) recvErr() (string, bool) {
	data, closed := self.recv()
	if closed {
		return "", true
	}

	if data == "PROJECT ID DNE" {
		return data, false
	}

	if strings.HasPrefix(data, "FAIL:") {
		return data, false
	}

	panic(data)
}

func connect(t *testing.T, userId string, projectId string) (TestConn, ProjectAndUsers) {
	a := make(chan string)
	b := make(chan string)
	closeCheck := make(chan struct{})

	forConnHandler := TestConn{tx: a, rx: b, closeCheck: closeCheck}
	forReturn := TestConn{tx: b, rx: a, closeCheck: closeCheck}

	go handleConnection(forConnHandler, userId)

	idMessage, disconnect := forReturn.recv()
	require.False(t, disconnect)
	require.Equal(t, idMessage, "WHOAMI: "+userId)

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

	time.Sleep(500 * time.Millisecond)

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

// USER STORY 2.7

func TestDeletingProject(t *testing.T) {
	resetDB()

	router := setupTestRouter()

	// Create an invite link
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/invite?id=53ed4d28-9279-4b4e-9256-b1e693332625", nil)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	count, _ := db.Collection("invitations").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)

	conn, _ := connect(t, "67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

	conn.send(`{
		"Name": "delete",
		"Args": {}
	}`)

	_, disconnect := conn.recvState()
	require.True(t, disconnect)

	user, err := getUserById(context.TODO(), "67c7b20021675682e4395270")
	require.Nil(t, err)
	require.Equal(t, len(user.Projects), 0)

	// Deleting doesn't care about invitations
	count, _ = db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)
}

func TestDeletingUserAround(t *testing.T) {
	resetDB()

	router := setupTestRouter()

	// Create an invite link
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/invite?id=53ed4d28-9279-4b4e-9256-b1e693332625", nil)
	router.ServeHTTP(w, req)
	bodyBytes, _ := io.ReadAll(w.Body)
	body := string(bodyBytes)
	require.Equal(t, http.StatusOK, w.Code)

	conn, project := connect(t, "67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

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
	project, disconnect := conn.recvState()
	require.False(t, disconnect)
	require.Equal(t, len(project.Users), 2)
	require.False(t, project.Users[0].LeftProject)
	require.False(t, project.Users[1].LeftProject)

	conn.send(`{
		"Name": "delete",
		"Args": {}
	}`)

	// There is another user in the project, so don't delete it
	err, _ := conn.recvErr()
	require.Equal(t, "FAIL: Cannot delete the project when there are other users still in the project", err)

	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

// PROJECT MUTATION ENDPOINTS

func TestProjectMutation(t *testing.T) {
	resetDB()

	conn, _ := connect(t, "67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

	conn.send(`{
		"Name": "update",
		"Args": {
			"Selector": "Title",
			"NewValue": "BRUH"
		}
	}`)

	proj, _ := conn.recvState()
	require.Equal(t, proj.Project.Title, "BRUH")

	// These endpoints shouldn't be able to touch `Users`
	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Users",
			"NewValue": {}
		}
	}`)

	_, _ = conn.recvErr()

	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Polls",
			"NewValue": {
				"Id": "poll",
				"Title": "Heyo",
				"Description": "Dank",
				"Options": [{
					"Id": "option 1",
					"Title": "Is Too",
					"LikedUsers": [],
					"NeutralUsers": [],
					"DislikedUsers": []
				}],
				"DueDate": 123
			}
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, len(proj.Project.Polls), 1)
	require.Equal(t, proj.Project.Polls[0].Id, "poll")
	require.Equal(t, proj.Project.Polls[0].DueDate, 123)
	require.Equal(t, len(proj.Project.Polls[0].Options), 1)
	require.Equal(t, proj.Project.Polls[0].Options[0].Id, "option 1")

	// Verify selector parsing
	conn.send(`{
		"Name": "update",
		"Args": {
			"Selector": "Polls[Id=poll].Options[Id=option 1].Title",
			"NewValue": "Is Not"
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, "Is Not", proj.Project.Polls[0].Options[0].Title)

	// ID doesn't exist, should fail silently
	conn.send(`{
		"Name": "update",
		"Args": {
			"Selector": "Polls[Id=polly].Options[Id=option 1].Title",
			"NewValue": "Is Not"
		}
	}`)

	_, _ = conn.recvState()

	// Struct field doesn't exist, should fail loudly
	conn.send(`{
		"Name": "update",
		"Args": {
			"Selector": "Polls[Id=poll].Optiony[Id=option 1].Title",
			"NewValue": "Is Not"
		}
	}`)

	_, _ = conn.recvErr()

	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Polls",
			"NewValue": {
				"Id": "poll2",
				"Title": "Heyo",
				"Description": "Dank",
				"Options": [],
				"DueDate": 123
			}
		}
	}`)

	_, _ = conn.recvState()

	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Polls",
			"NewValue": {
				"Id": "poll3",
				"Title": "Heyo",
				"Description": "Dank",
				"Options": [],
				"DueDate": 123
			}
		}
	}`)

	_, _ = conn.recvState()

	// Test with multiple options
	conn.send(`{
		"Name": "update",
		"Args": {
			"Selector": "Polls[Id=poll2].Title",
			"NewValue": "Heyoooo???"
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, "Heyoooo???", proj.Project.Polls[1].Title)

	conn.send(`{
		"Name": "update",
		"Args": {
			"Selector": "Polls[Id=poll3].Title",
			"NewValue": "Skibidi"
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, "Skibidi", proj.Project.Polls[2].Title)

	conn.send(`{
		"Name": "remove",
		"Args": {
			"Selector": "Polls[Id=poll2]"
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, 2, len(proj.Project.Polls))
	require.Equal(t, "Heyo", proj.Project.Polls[0].Title)
	require.Equal(t, "Skibidi", proj.Project.Polls[1].Title)

	// More deletion
	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Polls[Id=poll].Options[Id=option 1].LikedUsers",
			"NewValue": "Henry"
		}
	}`)

	_, _ = conn.recvState()

	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Polls[Id=poll].Options[Id=option 1].LikedUsers",
			"NewValue": "Ethan"
		}
	}`)

	_, _ = conn.recvState()

	conn.send(`{
		"Name": "append",
		"Args": {
			"Selector": "Polls[Id=poll].Options[Id=option 1].LikedUsers",
			"NewValue": "Cate"
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, proj.Project.Polls[0].Options[0].LikedUsers, []string{"Henry", "Ethan", "Cate"})

	conn.send(`{
		"Name": "remove",
		"Args": {
			"Selector": "Polls[Id=poll].Options[Id=option 1].LikedUsers[$IT=Ethan]"
		}
	}`)

	proj, _ = conn.recvState()

	require.Equal(t, proj.Project.Polls[0].Options[0].LikedUsers, []string{"Henry", "Cate"})

	time.Sleep(500 * time.Millisecond)
}
