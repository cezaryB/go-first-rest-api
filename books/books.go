package books

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var BOOKS = []book{
	{ Id: 1, Title: "Arkham horror", Author: "H.P.Lovecraft", Genre: horror },
	{ Id: 2, Title: "Lenore's shadow", Author: "E.A.Poe", Genre: horror },
	{ Id: 3, Title: "Shadow of the wind", Author: "Carlos Ruiz Zafon", Genre: thriller },
}

func HandleGetBooks(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	json.NewEncoder(w).Encode(BOOKS)
}

func HandleGetBook(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	idx := -1

	for i, v := range BOOKS {
		if fmt.Sprint(v.Id) == ps.ByName("id") {
			idx = i
		}
	}

	if idx == -1 {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(BOOKS[idx])
}

func HandlePostBook(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	t := req.FormValue("title")
	a := req.FormValue("author")
	g := req.FormValue("genre")
	genre, e := transformGenre(g)
	i := len(BOOKS) + 1

	if t == "" || a == "" || g == "" {
		http.Error(w, "You need to provide title, author and genre", http.StatusInternalServerError)
		return
	}

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	nBook := book{
		Id: i,
		Title: t,
		Author: a,
		Genre: genre,
	}

	BOOKS = append(BOOKS, nBook)	
	json.NewEncoder(w).Encode(nBook)
}

func HandleUpdateBook(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	t := req.FormValue("title")
	a := req.FormValue("author")
	g := req.FormValue("genre")
	genre, e := transformGenre(g)

	if e != nil && e.Error() != "Empty string" {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}

	b, e := findBookById(ps.ByName("id"), BOOKS)

	if e != nil {
		http.Error(w, e.Error(), http.StatusNotFound)
		return
	}

	if t != "" {
		b.Title = t
	}

	if a != "" {
		b.Author = a
	}

	if g != "" {
		b.Genre = genre
	}

	json.NewEncoder(w).Encode(b)
}
