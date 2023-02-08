package users

import (
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	bytePwd := []byte("example password")
	h, err := hashAndSalt(bytePwd)

	if h == string(bytePwd) {
		t.Errorf("Received hash was not encrypted")
	}

	if err != nil {
		t.Errorf("Encryption generation process failed")
	}
}

func TestVerifyIfPasswordsMatch(t *testing.T) {
	pwd1 := "password 1"
	h1, _ := hashAndSalt([]byte(pwd1))

	pwd2 := "password 2"
	h2, _ := hashAndSalt([]byte(pwd2))

	pwd3 := "password 3"
	h3, _ := hashAndSalt([]byte(pwd3))

	tables := []struct{
		p string
		h string
		o bool
	}{
		{pwd1, h1, true},
		{pwd2, h2, true},
		{"incorrect password", h3, false},
	}

	for _, table := range tables {
		match := verifyIfPasswordsMatch(table.h, table.p)

		if match != table.o {
			t.Errorf("Verification process failed, expected %v, got %v", table.o, match)
		}
	}
}

func TestValidateUsername(t *testing.T) {
	users := []user{
		{ Username: "Username1", Email: "email@one", password: "password1" },
		{ Username: "Username2", Email: "email@two", password: "password2" },
	}

	tables := []struct{
		u string
		o bool
	}{
		{ "Username1", false},
		{ "Username2", false},
		{ "Username3", true},
	}

	for _, table := range tables {
		v := validateUsername(table.u, users)

		if v != table.o {
			t.Errorf("Validation process failed, expected %v, got %v", table.o, v)
		}
	}
}