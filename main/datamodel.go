package main

type AuthorDto struct {
	UUID      string
	FirstName string
	Lastname  string
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
