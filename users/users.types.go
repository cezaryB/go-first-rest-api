package users

type user struct {
	Username string
	Email string
	password string
}

type authStatus struct {
	Authenticated bool
}

type createUserDTO struct {
	Username string
	Email string
	Password string
}

type authenticateUserDTO struct {
	Username string
	Password string
}
