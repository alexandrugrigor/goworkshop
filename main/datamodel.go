package main

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string
	Title       string
	NoPages     int
	ReleaseDate string
	Author      AuthorDto
}

//AuthorDto - The DTO used to access authors
type AuthorDto struct {
	UUID      string
	FirstName string
	LastName  string
	Birthday  string
	Death     string
}

//Books - the list of available books
var Books = []BookDto{}

// Authors - the list of available authors
var Authors = []AuthorDto{}
