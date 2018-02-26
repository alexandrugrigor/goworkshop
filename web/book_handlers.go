package web

import (
	"net/http"
	"github.com/gorilla/mux"
	"goworkshop/model"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

//Demonstrates the basic functionality of private and public modifiers in GO
func Index(w http.ResponseWriter, r *http.Request) {
	helloWorkshop := struct {
		Message        string `json:"message"`
		privateMessage string `json:"privateMessage"`
		NoTagField     string `json:"-"`
	}{
		Message:        "Hello workshop!",
		privateMessage: "Message that does not appear in response :).",
		NoTagField:     "This message won't appear either",
	}
	WriteJson(w, helloWorkshop)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, model.Books)
}

func GetBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	book, err := model.Books.Get(bookUUID)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, book)
	}
}

func DeleteBookByUUID(w http.ResponseWriter, r *http.Request) {
	var bookUUID = mux.Vars(r)["uuid"]
	err := model.Books.Delete(bookUUID)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete book: %s", err)
	} else {
		WriteJson(w, model.Books)
	}
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.BookDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to create book: %s", err)
	} else {
		model.Books.Add(book)
		WriteJson(w, book)
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.BookDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &book)
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}
	book, err = model.Books.Update(book)
	if err != nil {
		fmt.Fprintf(w, "Failed to update book: %s", err)
		return
	}
	WriteJson(w, book)
}