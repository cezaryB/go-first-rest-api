package books

type genre int

const (
	horror genre = iota
	comedy
	thriller
	criminal
	scifi
	fantasy
	none
)

type book struct {
	Id int
	Title string
	Author string
	Genre genre
}

type createBookDTO struct {
	Title string
	Author string
	Genre string
}

type updateBookDTO = createBookDTO