module main

replace books => ./books

replace middleware => ./middleware

go 1.19

require (
	books v0.0.0-00010101000000-000000000000 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	middleware v0.0.0-00010101000000-000000000000 // indirect
)
