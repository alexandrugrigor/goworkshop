package importer

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
	"goworkshop/model"
)

func ImportAuthors() []model.AuthorDto {
	fileContent, err := ioutil.ReadFile("importer/authors.json")
	if err != nil {
		fmt.Printf("Failed to read authors.json file: %s\n", err)
		os.Exit(1)
	}
	var authors []model.AuthorDto
	err = json.Unmarshal(fileContent, &authors)
	if err != nil {
		fmt.Printf("Failed to unmarshal authors: %s\n", err)
		os.Exit(1)
	}
	return authors
}

func ImportBooks() []model.BookDto {
	fileContent, err := ioutil.ReadFile("importer/books.json")
	if err != nil {
		fmt.Printf("Failed to read book.json file: %s\n", err)
		os.Exit(1)
	}
	var books []model.BookDto
	err = json.Unmarshal(fileContent, &books)
	if err != nil {
		fmt.Printf("Failed to unmarshal books: %s\n", err)
		os.Exit(1)
	}
	return books
}
