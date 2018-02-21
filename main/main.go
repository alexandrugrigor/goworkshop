package main

import (
	"fmt"
	"goworkshop/model"
	"goworkshop/test"
)

func main() {
	var s = test.CreateSquare(11)
	s.Color = "green"
	fmt.Println(s.Area())
	fmt.Println(s)
	var book = model.AuthorDto{
		UUID:      "1234",
		FirstName: "Ion",
		LastName:  "Creanga",
		Birthday:  "1850",
		Death:     "1914",
	}

	fmt.Println(book)

}
