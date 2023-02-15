package users

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var USERS = []user{
	{Username: "Example user", Email: "example@gmail.com", password: "$2a$04$WxRvBX5FDN6xanjSSg73Tu6V6yBlYvTomSRec6gWL/WHUvK9SEzNu"},
}

func HandlePostUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	var payload createUserDTO
	json.NewDecoder(req.Body).Decode(&payload)
	u, e, p := payload.Username, payload.Email, payload.Password

	if u == "" || p == "" || e == "" {
		http.Error(w, "You need to provide username, email and password", http.StatusBadRequest)
		return
	}

	if !validateUsername(u, USERS) {
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
		Username: u,
		Email: e,
		password: string(hashedPwd),
	}

	USERS = append(USERS, nUser)
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(nUser)
}

func HandleAuthenticateUser(w http.ResponseWriter, req * http.Request, _ httprouter.Params) {
	var user user
	var payload authenticateUserDTO

	json.NewDecoder(req.Body).Decode(&payload)
	u, p := payload.Username, payload.Password
	s := authStatus{ Authenticated: false }

	if u == "" || p == "" {
		http.Error(w, "You need to provide username and password", http.StatusInternalServerError)
		return
	}

	for _, v := range USERS {
		if v.Username == u {
			user = v
		}
	}

	if verifyIfPasswordsMatch(user.password, p) {
		s = authStatus{ Authenticated: true }
	}

	if !s.Authenticated	{
		w.WriteHeader(http.StatusUnauthorized)
	}

	json.NewEncoder(w).Encode(s)
}