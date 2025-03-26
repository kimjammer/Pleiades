package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// I don't know how to test with a real websocket connection so this will have to do. The playwright tests will probably implicitly test that anyways.
func testApply(projectId string, userId string, message string) error {
	command := CommandMessage{}

	err := json.Unmarshal(([]byte)(message), &command)
	if err != nil {
		panic(err)
	}

	decodedCommand, err := decodeCommand(command, userId)
	if err != nil {
		panic(err)
	}

	return applyCommandToProject(projectId, decodedCommand)
}

// USER STORY 2.6

func TestLeavingDeletesProject(t *testing.T) {
	resetDB()

	leave("67c7b20021675682e4395270", "53ed4d28-9279-4b4e-9256-b1e693332625")

	user, err := getUserById(context.TODO(), "67c7b20021675682e4395270")
	require.Nil(t, err)
	require.Equal(t, len(user.Projects), 0)

	require.Equal(t, testApply("53ed4d28-9279-4b4e-9256-b1e693332625", "67c7b20021675682e4395270", `{
		"Name": "leave",
		"Args": {}
	}`), nil)

	// With no invite links around, leaving a project as the only user should delete it
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)
}
