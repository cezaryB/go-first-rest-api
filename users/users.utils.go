package users

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(bytePwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)

	if err != nil {
		log.Println(err)
	}

	return string(hash), err
}

func verifyIfPasswordsMatch(hashedPwd string, plainPwd string) bool {
	byteHashedPwd := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHashedPwd, bytePlainPwd)

	if err != nil {
		return false
	}

	return true
}

func validateUsername(username string, users []user) bool {
	for _, v := range users {
		if v.Username == username {
			return false
		}
	}

	return true
}
