package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewProjectNoAuth(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", nil)
	router.ServeHTTP(w, req)

	//Status Code is 401 Unauthorized
	require.Equal(t, http.StatusUnauthorized, w.Code)

	//DB is not modified
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

func TestNewProjectNoBody(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	//DB is not modified
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

func TestNewProjectEmptyTitle(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	request := NewProjectRequest{
		Title:       "",
		Description: "",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)

	//DB is not modified
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

func TestNewProjectSuccess(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	request := NewProjectRequest{
		Title:       "Test Project",
		Description: "Test Description",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/projects/new", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(2), count)
}

func TestGetProjectsNoAuth(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/projects", nil)
	router.ServeHTTP(w, req)

	//Status Code is 401 Unauthorized
	require.Equal(t, http.StatusUnauthorized, w.Code)

	//DB is not modified
	count, _ := db.Collection("projects").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

func TestGetProjects(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/projects", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err.Error())
	}

	var data ProjectsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}
	log.Println(data)

	require.Equal(t, 1, len(data.Projects))
}
