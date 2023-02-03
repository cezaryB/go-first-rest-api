package main

import (
	b "books"
	"log"
	m "middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)



func main() {
	router := httprouter.New()

	router.GET("/books", m.JSONcontentMidleware(b.HandleGetBooks))
	router.GET("/books/:id", m.JSONcontentMidleware(b.HandleGetBook))
	router.POST("/books", m.JSONcontentMidleware(b.HandlePostBook))
	router.PATCH("/books/:id", m.JSONcontentMidleware(b.HandleUpdateBook))

	log.Fatal(http.ListenAndServe(":8080", router))
	
}

