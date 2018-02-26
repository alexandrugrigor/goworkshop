package web

import (
	"net/http"
	"goworkshop/importer"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
)


type Authors []model.AuthorDto

var authors Authors = importer.ImportAuthors()

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, authors)
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	var authorUUID = mux.Vars(r)["uuid"]
	author, err := authors.get(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, author)
	}
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	var authorUUID = mux.Vars(r)["uuid"]
	err := authors.delete(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete author: %s", err)
	} else {
		WriteJson(w, authors)
	}
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.AuthorDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to create author: %s", err)
	} else {
		authors = append(authors, author)
		WriteJson(w, author)
	}
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.AuthorDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	author, err = authors.update(author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	WriteJson(w, author)
}

func (a *Authors) get(authorUUID string) (model.AuthorDto, error) {
	err := fmt.Errorf("could not find author by uuid %s", authorUUID)
	for _, author := range *a {
		if author.UUID == authorUUID {
			return author, nil
		}
	}
	var author model.AuthorDto
	return author, err
}

func (a *Authors) delete(authorUUID string) error {
	var err error = fmt.Errorf("could not find author by uuid %s", authorUUID)
	var updatedAuthors Authors
	for _, author := range *a {
		if author.UUID == authorUUID {
			err = nil
		} else {
			updatedAuthors = append(updatedAuthors, author)
		}
	}
	if err == nil {
		*a = updatedAuthors
	}
	return err
}

func (a *Authors) update(updatedAuthor model.AuthorDto) (model.AuthorDto, error) {
	var err error = fmt.Errorf("could not find author by uuid %s", updatedAuthor.UUID)
	var newAuthors Authors
	for _, author := range *a {
		if author.UUID == updatedAuthor.UUID {
			newAuthors = append(newAuthors, updatedAuthor)
			err = nil
		} else {
			newAuthors = append(newAuthors, author)
		}
	}
	if err == nil {
		*a = newAuthors
	}
	return updatedAuthor, err
}