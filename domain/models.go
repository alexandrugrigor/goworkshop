package domain

type AuthorDto struct {
	UUID      string
	FirstName string
	LastName  string
	Birthday  string
	Death     string
}

type BookDto struct {
	UUID        string
	Title       string
	NoPages     int
	ReleaseDate string
	Author      AuthorDto
}
