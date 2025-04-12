package main

import (
	"bytes"
	"context"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
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

func TestVerifySessionNoAuth(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/verifySession", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestVerifySession(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/verifySession", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

// User story 2.2 acceptancance criteria 1
func TestInviteCreation(t *testing.T) {
	router := setupTestRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/invite?id=53ed4d28-9279-4b4e-9256-b1e693332625", nil)
	router.ServeHTTP(w, req)
	// w.Body = uuid
	// Validation tested in e2e

	require.Equal(t, http.StatusOK, w.Code)

	//DB is modified
	count, _ := db.Collection("invitations").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

// User story 2.2 acceptancance criteria 2
func TestInviteUnauthorized(t *testing.T) {
	// TODO: very similar to `TestGetProjectsNoAuth` make sense to abstract?

	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/invite?id=53ed4d28-9279-4b4e-9256-b1e693332625", nil)
	router.ServeHTTP(w, req)

	//Status Code is 401 Unauthorized
	require.Equal(t, http.StatusUnauthorized, w.Code)

	//DB is not modified
	count, _ := db.Collection("invitations").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)
}

func TestCheckEmailUsed(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/register/check?email=example%40example.com", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err.Error())
	}

	var data EmailCheckResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	require.Equal(t, true, data.Exists)
}

func TestCheckEmailNotUsed(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/register/check?email=notused%40example.com", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err.Error())
	}

	var data EmailCheckResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	require.Equal(t, false, data.Exists)
}

func TestRegisterUser(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	request := User{
		FirstName: "Steve",
		LastName:  "Minecraft",
		Email:     "notch@minecraft.net",
		Password:  "blockgames4ever",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/register/new", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusCreated, w.Code)

	count, _ := db.Collection("users").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(2), count)
}

func TestRegisterUserInvalid(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	jsonData, _ := json.Marshal("invalid garbage")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/register/new", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/login?email=example%40example.com&password=notthepassword", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err.Error())
	}

	var data LoginResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	require.Equal(t, true, data.Exists)
}

func TestLoginUserInvalidPassword(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/login?email=example%40example.com&password=invalidpassword", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err.Error())
	}

	var data LoginResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	require.Equal(t, false, data.Exists)
}

func TestLoginInvalidUser(t *testing.T) {
	router := setupRouter()
	defineRoutes(router)
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/login?email=invalid%40example.com&password=doesntmatter", nil)
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body, err := io.ReadAll(w.Body)
	if err != nil {
		panic(err.Error())
	}

	var data LoginResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err.Error())
	}

	require.Equal(t, false, data.Exists)
}

func TestLogout(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/logout", nil)
	router.ServeHTTP(w, req)
	require.Equal(t, w.Result().Cookies()[0].SameSite, http.SameSiteNoneMode)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestStats(t *testing.T) {
	// Setup
	router := setupTestRouter()
	resetDB()
	w := httptest.NewRecorder()
	var resp map[string]any

	// Is empty on first load
	req, _ := http.NewRequest(http.MethodGet, "/stats", nil)
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Empty(t, resp)

	// Increment stat by one
	req = httptest.NewRequest(http.MethodGet, "/event?name=test", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	// Check new state
	req = httptest.NewRequest(http.MethodGet, "/stats", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, float64(1), resp["test"])

	// Increment stat by ten
	req = httptest.NewRequest(http.MethodGet, "/event?name=test&value=10", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	require.Equal(t, http.StatusOK, w.Code)

	// Check new state
	req = httptest.NewRequest(http.MethodGet, "/stats", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, float64(11), resp["test"])
}

func TestForgotPasswordInvalidEmail(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	request := ForgotPasswordRequest{
		Email: "invalid@email.com",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/forgotPassword", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	//Status should be 200 OK since we don't want bad actors to know if the email is valid
	require.Equal(t, http.StatusOK, w.Code)

	//No tokens should be created
	count, _ := db.Collection("pwdResetTokens").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)
}

func TestForgotPasswordValidEmail(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	request := ForgotPasswordRequest{
		Email: "example@example.com",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/forgotPassword", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	//1 token should be created
	count, _ := db.Collection("pwdResetTokens").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(1), count)
}

func TestResetPasswordInvalidToken(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	request := PasswordResetRequest{
		Token:       "invalid",
		NewPassword: "doesn't matter",
	}
	jsonData, _ := json.Marshal(request)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/resetPassword", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestResetPasswordValidToken(t *testing.T) {
	router := setupTestRouter()
	resetDB()

	//Create token first
	forgotRequest := ForgotPasswordRequest{
		Email: "example@example.com",
	}
	jsonData, _ := json.Marshal(forgotRequest)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/forgotPassword", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	//Get token from DB
	var pwdResetToken PwdResetToken
	_ = db.Collection("pwdResetTokens").FindOne(context.TODO(), bson.D{}).Decode(&pwdResetToken)

	//Send reset request
	resetRequest := PasswordResetRequest{
		Token:       pwdResetToken.Token,
		NewPassword: "newpassword",
	}
	jsonData, _ = json.Marshal(resetRequest)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/resetPassword", bytes.NewBuffer(jsonData))
	router.ServeHTTP(w, req)

	//Check status
	require.Equal(t, http.StatusOK, w.Code)

	//Check token is now deleted
	count, _ := db.Collection("pwdResetTokens").CountDocuments(context.TODO(), bson.D{})
	require.Equal(t, int64(0), count)

	//Check password is updated
	var user User
	_ = db.Collection("users").FindOne(context.TODO(), bson.D{}).Decode(&user)
	result := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("newpassword"))
	require.Nil(t, result)
}
