package main

import (
	"fmt"
	"goworkshop/importer"
	"goworkshop/model"
	"goworkshop/web"
	"goworkshop/persistence"
)

func main() {
	persistence.InitDB()
	model.Authors = importer.ImportAuthors()
	fmt.Printf("Imported authors are: %s\n", model.Authors)
	model.Books = importer.ImportBooks()
	fmt.Printf("Imported books are: %s\n", model.Books)
	web.StartServer()
}

