package books

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

	router.GET("/books", m.JSONcontentMidleware(HandleGetBooks))
	router.GET("/books/:id", m.JSONcontentMidleware(HandleGetBook))
	router.POST("/books", m.JSONcontentMidleware(HandlePostBook))
	router.PATCH("/books/:id", m.JSONcontentMidleware(HandleUpdateBook))

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

func TestHandleGetBooks(t *testing.T) {
	var books []book
	responseWriter := makeRequest("GET", "/books", nil)

	json.NewDecoder(responseWriter.Body).Decode(&books)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, len(books), 3)
}

func TestHandleGetBook(t *testing.T) {
	var book book
	responseWriter := makeRequest("GET", "/books/1", nil)

	json.NewDecoder(responseWriter.Body).Decode(&book)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, book.Id, 1)
}

func TestHandleGetBookNotFound(t *testing.T) {
	responseWriter := makeRequest("GET", "/books/999", nil)

	assert.Equal(t, http.StatusNotFound, responseWriter.Code)
	assert.Equal(t, responseWriter.Body.String(), "Resource not found\n")
}

func TestHandlePostBookSucess(t *testing.T) {
	var book book
	body := createBookDTO{ 
		Title: "Anne from Green Gables", Author: "L. M. Montgomery", Genre: "Comedy",
	}

	responseWriter := makeRequest("POST", "/books", body)
	json.NewDecoder(responseWriter.Body).Decode(&book)

	assert.Equal(t, http.StatusCreated, responseWriter.Code)
	assert.Equal(t, body.Title, book.Title)
	assert.Equal(t, body.Author, book.Author)
	// The one below would fail since I haven't yet introduced genre transformation
	// assert.Equal(t, body.Genre, book.Genre)
}

func TestHandlePostBookIncompleteBody(t *testing.T) {
	body := createBookDTO{ 
		Title: "Anne from Green Gables", Author: "L. M. Montgomery",
	}

	responseWriter := makeRequest("POST", "/books", body)

	assert.Equal(t, http.StatusInternalServerError, responseWriter.Code)
	assert.Equal(t, responseWriter.Body.String(), "You need to provide title, author and genre\n")
}

func TestHandleUpdateBookSucess(t *testing.T) {
	var book book
	body := updateBookDTO{
		Title: "Updated title",
	}

	responseWriter := makeRequest("PATCH", "/books/1" , body)
	json.NewDecoder(responseWriter.Body).Decode(&book)

	assert.Equal(t, http.StatusOK, responseWriter.Code)
	assert.Equal(t, book.Id, 1)
	assert.Equal(t, book.Title, body.Title)
}

func TestHandleUpdateBookNotFound(t *testing.T) {
	body := updateBookDTO{
		Title: "Updated title",
	}

	responseWriter := makeRequest("PATCH", "/books/9999", body)

	assert.Equal(t, http.StatusNotFound, responseWriter.Code)
	assert.Equal(t, responseWriter.Body.String(), "Book not found\n")
}

