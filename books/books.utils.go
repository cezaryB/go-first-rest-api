package books

import (
	"errors"
	"strings"
)

func TransformGenre(g string) (Genre, error) {
	var gtransformed Genre
	var err error
	g = strings.ToUpper(g)
	
	switch g {
	case "HORROR":
		gtransformed = Horror
	case "COMEDY":
		gtransformed = Comedy
	case "THRILLER":
		gtransformed = Thriller
	case "CRIMINAL":
		gtransformed = Criminal
	case "SciFi":
		gtransformed = SciFi
	case "FANTASY":
		gtransformed = Fantasy
	case "":
		err = errors.New("Empty string")
	default:
		err = errors.New("Genre has to be one of the following: 'Horror', 'Comedy', 'Thriller', 'Criminal', 'SciFi', 'Fantasy'")
	}

	return gtransformed, err
}