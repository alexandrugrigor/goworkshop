package model

import (
	"fmt"
)

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}

//AuthorDto - The DTO used to access authors
type AuthorDto struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

func (author AuthorDto) String() string {
	return fmt.Sprintf("AuthorDto{UUID='%s', FirstName='%s', LastName='%s', Birthday='%s', Death='%s'}", author.UUID,
		author.FirstName, author.LastName, author.Birthday, author.Death)
}

func (book BookDto) String() string {
	return fmt.Sprintf("BookDto{UUID='%s', Title='%s', NoPages=%d, ReleaseDate='%s',Author=%s}", book.UUID, book.Title, book.NoPages, book.ReleaseDate, book.Author)
}

//Books - the list of available books
var Books []BookDto

// Authors - the list of available authors
var Authors []AuthorDto
