package books

import (
	"errors"
	"reflect"
	"testing"
)

func TestTransformGenre(t *testing.T) {
	var e error
	tables := []struct{
		x string
		n genre
		e error
	}{
		{"horror", horror, e},
		{"comedy", comedy, e},
		{"thriller", thriller, e},
		{"criminal", criminal, e},
		{"scifi", scifi, e},
		{"fantasy", fantasy, e},
		{"invalid genre", none, errors.New("Genre has to be one of the following: 'Horror', 'Comedy', 'Thriller', 'Criminal', 'SciFi', 'Fantasy'")},
		{"", none, errors.New("Empty string")},
	}

	for _, table := range tables {
		g, err := transformGenre(table.x)

		if g != table.n {
			t.Errorf("Transformed genre was incorrect, expected %v, got %v", table.n, g)
		}

		if err != nil && err.Error() != table.e.Error() {
			t.Errorf("Error message was incorrect, expected %v, got %v", table.e.Error(), err.Error())
		}
	}
}

func TestFindBookById(t *testing.T) {
	var e error
	var eBook book
	books := []book{
		{ Id: 1, Title: "Arkham horror", Author: "H.P.Lovecraft", Genre: horror },
		{ Id: 2, Title: "Lenore's shadow", Author: "E.A.Poe", Genre: horror },
	}
	tables := []struct{
		x string
		b book
		e error
	}{
		{ "1", books[0], e},
		{ "2", books[1], e},
		{ "3", eBook, errors.New("Book not found")},
	}

	for _, table := range tables {
		b, err := findBookById(table.x, books)

		if err == nil && !reflect.DeepEqual(b, table.b) {
			t.Errorf("Retuned book was incorrect, expected %v, got %v", table.b, b)
		}

		if err != nil && err.Error() != table.e.Error() {
			t.Errorf("Error message was incorrect, expected %v, got %v", table.e.Error(), err.Error())
		}
	}

}