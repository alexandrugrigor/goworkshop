package model

//AuthorDto - Author structure
type AuthorDto struct {
	UUID      string
	FirstName string
	LastName  string
	Birthday  string
	Death     string
}

//BookDto - Book structure
type BookDto struct {
	UUID        string
	Title       string
	NoPages     int
	ReleaseDate string
	Author      AuthorDto
}

var authors []AuthorDto
var books []BookDto
