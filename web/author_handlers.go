package web

import (
	"net/http"
	"goworkshop/model"
	"github.com/gorilla/mux"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	WriteJson(w, model.Authors)
}

func GetAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	authorUUID := mux.Vars(r)["uuid"]
	author, err := model.Authors.Get(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	} else {
		WriteJson(w, author)
	}
}

func DeleteAuthorByUUID(w http.ResponseWriter, r *http.Request) {
	var authorUUID = mux.Vars(r)["uuid"]
	err := model.Authors.Delete(authorUUID)
	if err != nil {
		fmt.Fprintf(w, "Failed to delete author: %s", err)
	} else {
		WriteJson(w, model.Authors)
	}
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.AuthorDto
	bytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(bytes, &author)
	if err != nil {
		fmt.Fprintf(w, "Failed to create author: %s", err)
	} else {
		model.Authors.Add(author)
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
	author, err = model.Authors.Update(author)
	if err != nil {
		fmt.Fprintf(w, "Failed to update author: %s", err)
		return
	}
	WriteJson(w, author)
}
