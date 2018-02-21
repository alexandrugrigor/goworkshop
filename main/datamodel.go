package main

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
	NoPages     string
	ReleaseDate string
	Author      AuthorDto
}
