package books

import (
	"testing"
)

func TestTransformGenre(t * testing.T) {
	g, err := transformGenre("")

	if err.Error() != "Empty string" {
		t.Errorf("Error message was incorrect, expect 'Empty String', got %v", err.Error())
	}

	if g != 0 {
		t.Errorf("Genre assigned by default was incorrect, expected 0, got %v", g)
	}
}