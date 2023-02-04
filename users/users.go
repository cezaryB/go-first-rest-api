package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var USERS = []user{}

func HandlePostUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	n := req.FormValue("username")
	e := req.FormValue("email")
	p := req.FormValue("password")

	if n == "" || p == "" || e == "" {
		http.Error(w, "You need to provide username, email and password", http.StatusInternalServerError)
		return
	}

	if !verifyIfUsernameIsValid(n, USERS) {
		http.Error(w, "This username is already taken", http.StatusBadRequest)
		return
	}

	bytePwd := []byte(p)
	hashedPwd, err := hashAndSalt(bytePwd)

	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	nUser := user{
		Username: n,
		Email: e,
		password: string(hashedPwd),
	}

	USERS = append(USERS, nUser)
	fmt.Println(USERS)
	json.NewEncoder(w).Encode(nUser)
}

func HandleAuthenticateUser(w http.ResponseWriter, req * http.Request, _ httprouter.Params) {
	n := req.FormValue("username")
	p := req.FormValue("password")
	s := authStatus{ Authenticated: false }
	var u user

	if n == "" || p == "" {
		http.Error(w, "You need to provide username and password", http.StatusInternalServerError)
		return
	}

	for _, v := range USERS {
		if v.Username == n {
			u = v
		}
	}

	if verifyIfPasswordsMatch(u.password, p) {
		s = authStatus{ Authenticated: true }
	}

	json.NewEncoder(w).Encode(s)
}