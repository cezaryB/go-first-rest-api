package books

type genre int

const (
	horror genre = iota + 1
	comedy
	thriller
	criminal
	scifi
	fantasy
)

type book struct {
	Id int
	Title string
	Author string
	Genre genre
}