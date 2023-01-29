package books

type Genre int

const (
	Horror Genre = iota
	Comedy
	Thriller
	Criminal
	SciFi
	Fantasy
)

type book struct {
	Id int
	Title string
	Author string
	Genre Genre
}