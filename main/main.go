package main

import (
	"fmt"
	"goworkshop/test"
)

func main() {
	var s = test.Square{
		Rectangle: test.Rectangle{
			Length: 10, Width: 10,
		},

		Color: "red",
	}

	fmt.Println(s)

	var aut = AuthorDto{
		UUID:      "aut1",
		FirstName: "fs",
		LastName:  "ls",
		Birthday:  "1997",
		Death:     "15"}
	fmt.Println(aut)

}

//Square-square struct
// S - Upper case means public
// s-lower case mean package-private

type Square struct {
	Length int
	Color  string
}
