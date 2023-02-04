package users

type user struct {
	Username string
	Email string
	password string
}

type authStatus struct {
	Authenticated bool
}