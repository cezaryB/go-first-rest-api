package books

import (
	"errors"
	"fmt"
	"strings"
)

func transformGenre(g string) (genre, error) {
	var gtransformed genre
	var err error
	g = strings.ToUpper(g)
	
	switch g {
	case "HORROR":
		gtransformed = horror
	case "COMEDY":
		gtransformed = comedy
	case "THRILLER":
		gtransformed = thriller
	case "CRIMINAL":
		gtransformed = criminal
	case "SciFi":
		gtransformed = scifi
	case "FANTASY":
		gtransformed = fantasy
	case "":
		err = errors.New("Empty string")
	default:
		err = errors.New("Genre has to be one of the following: 'Horror', 'Comedy', 'Thriller', 'Criminal', 'SciFi', 'Fantasy'")
	}

	return gtransformed, err
}

func findBookById(id string, books []book) (book, error) {
	var err error
	var b book
	idx := -1

	for i, v := range books {
		if fmt.Sprint(v.Id) == id {
			idx = i
		}
	}

	if idx > -1 {
		return books[idx], err
	}

	err = errors.New("Book not found")

	return b, err
}