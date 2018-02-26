package model

import "fmt"

//Books - the list of available books
var Books BooksList

//BookDto - The DTO used to access books
type BookDto struct {
	UUID        string    `json:"uuid"`
	Title       string    `json:"title"`
	NoPages     int       `json:"noPages"`
	ReleaseDate string    `json:"releaseDate"`
	Author      AuthorDto `json:"author"`
}

func (book BookDto) String() string {
	return fmt.Sprintf("BookDto{UUID='%s', Title='%s', NoPages=%d, ReleaseDate='%s',Author=%s}",
		book.UUID, book.Title, book.NoPages, book.ReleaseDate, book.Author)
}

//BooksList - A list of BookDto objects
type BooksList []BookDto

func (b *BooksList) Get(bookUUID string) (BookDto, error) {
	err := fmt.Errorf("could not find book by uuid %s", bookUUID)
	for _, book := range *b {
		if book.UUID == bookUUID {
			return book, nil
		}
	}
	return BookDto{}, err
}

func (b *BooksList) Delete(bookUUID string) error {
	var err = fmt.Errorf("could not find book by uuid %s", bookUUID)
	var updatedBooks BooksList
	for _, book := range *b {
		if book.UUID == bookUUID {
			err = nil
		} else {
			updatedBooks = append(updatedBooks, book)
		}
	}
	if err == nil {
		*b = updatedBooks
	}
	return err
}

func (b *BooksList) Update(updatedBook BookDto) (BookDto, error) {
	var err = fmt.Errorf("could not find book by uuid %s", updatedBook.UUID)
	var newBooks BooksList
	for _, book := range *b {
		if book.UUID == updatedBook.UUID {
			newBooks = append(newBooks, updatedBook)
			err = nil
		} else {
			newBooks = append(newBooks, book)
		}
	}
	if err == nil {
		*b = newBooks
	}
	return updatedBook, err
}


//Adds the given book into the list
func (b *BooksList) Add(book BookDto){
	*b = append(*b, book)
}