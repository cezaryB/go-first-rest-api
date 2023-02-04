package main

import (
	b "books"
	"log"
	m "middleware"
	"net/http"
	u "users"

	"github.com/julienschmidt/httprouter"
)



func main() {
	p := ":8080"
	router := httprouter.New()

	router.GET("/books", m.JSONcontentMidleware(b.HandleGetBooks))
	router.GET("/books/:id", m.JSONcontentMidleware(b.HandleGetBook))
	router.POST("/books", m.JSONcontentMidleware(b.HandlePostBook))
	router.PATCH("/books/:id", m.JSONcontentMidleware(b.HandleUpdateBook))

	router.POST("/users", m.JSONcontentMidleware(u.HandlePostUser))
	router.POST("/users/login", m.JSONcontentMidleware(u.HandleAuthenticateUser))

	log.Fatal(http.ListenAndServe(p, m.RouterWithLog(router, p)))
	
}

