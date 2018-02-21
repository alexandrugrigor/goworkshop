package main

// AuthorDto will store the information about the authors
type AuthorDto struct {
	UUID      string
	FirstName string
	LastName  string
	Birthday  string
	Death     string
}

// BookDto will store the information about the books
type BookDto struct {
	UUID        string
	Title       string
	NoPages     int
	ReleaseDate string
	Author      AuthorDto
}
