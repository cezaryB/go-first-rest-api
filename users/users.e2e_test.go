package users

import (
	"bytes"
	"encoding/json"
	m "first-rest-api/middleware"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func router() *httprouter.Router {
	router := httprouter.New()

	router.POST("/users", m.JSONcontentMidleware(HandlePostUser))
	router.POST("/users/login", m.JSONcontentMidleware(HandleAuthenticateUser))

	return router
}

func makeRequest(method string, url string, body interface{}) *httptest.ResponseRecorder {
	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	responseWriter := httptest.NewRecorder()
	router().ServeHTTP(responseWriter, request)

	return responseWriter
}

func TestHandlePostUserSuccess(t *testing.T) {
	var user user
	body := createUserDTO{
		Username: "James",
		Email: "james@gmail.com",
		Password: "example123",
	}

	responseWriter := makeRequest("POST", "/users", body)
	json.NewDecoder(responseWriter.Body).Decode(&user)

	assert.Equal(t, http.StatusCreated, responseWriter.Code)
	assert.Equal(t, body.Username, user.Username)
	assert.Equal(t, body.Email, user.Email)
}

func TestHandlePostUserDuplicatedUsername(t *testing.T) {
	body := createUserDTO{
		Username: "Example user",
		Email: "example@gmail.com",
		Password: "example123",
	}

	responseWriter := makeRequest("POST", "/users", body)

	assert.Equal(t, http.StatusBadRequest, responseWriter.Code)
	assert.Equal(t, responseWriter.Body.String(), "This username is already taken\n")
}

func TestHandlePostUserIncompleteBody(t *testing.T) {
	body := createUserDTO{
		Username: "Example user",
		Email: "",
		Password: "example123",
	}

	responseWriter := makeRequest("POST", "/users", body)

	assert.Equal(t, http.StatusBadRequest, responseWriter.Code)
	assert.Equal(t, responseWriter.Body.String(), "You need to provide username, email and password\n")
}

func TestHandleAuthenticateUserSuccess(t *testing.T) {
	var status authStatus
	body := authenticateUserDTO{
		Username: "Example user",
		Password: "password123",
	}

	responseWriter := makeRequest("POST", "/users/login", body)
	json.NewDecoder(responseWriter.Body).Decode(&status)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, status.Authenticated, true)
}

func TestHandleAuthenticateUserFail(t *testing.T) {
	var status authStatus
	body := authenticateUserDTO{
		Username: "Example user",
		Password: "password345",
	}

	responseWriter := makeRequest("POST", "/users/login", body)
	json.NewDecoder(responseWriter.Body).Decode(&status)

	assert.Equal(t, http.StatusUnauthorized, responseWriter.Code)
	assert.Equal(t, status.Authenticated, false)
}

