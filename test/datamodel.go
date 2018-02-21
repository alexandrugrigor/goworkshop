// //create AuthorDto struct with the following fields:
// UUID string
// FirstName string
// LastName string
// Birthday string
// Death string

// create BookDto struct with the following fields:
// UUID string
// Title string
// NoPages int
// ReleaseDate string
// Author AuthorDto

package main

type AuthorDto struct {
	UUID, FirstName, LastName, Birthday, Death string
}

type BookDto struct {
	UUID, Title, ReleaseDate string
	NoPages                  int
	Author                   AuthorDto
}
