module main

replace books => ./books

replace middleware => ./middleware

replace users => ./users

go 1.19

require (
	books v0.0.0-00010101000000-000000000000
	github.com/julienschmidt/httprouter v1.3.0
	middleware v0.0.0-00010101000000-000000000000
	users v0.0.0-00010101000000-000000000000
)

require golang.org/x/crypto v0.5.0 // indirect
