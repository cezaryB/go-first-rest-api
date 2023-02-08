package books

import (
	"errors"
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
			t.Errorf("Error message was incorrect, expected %v, got %v", err.Error(), table.e.Error())
		}
	}
}