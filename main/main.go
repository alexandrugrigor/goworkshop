package main

import "fmt"

import "github.com/rusucosmin/goworkshop/test"

func main() {
	var books = make([]BookDto, 8)
	var authors = make([]AuthorDto, 8)

	fmt.Printf("%d\n", len(books))
	fmt.Printf("%d\n", len(authors))

	var v = test.Square{
		Length: 25,
	}

	var rect = test.Rectangle{
		Square: v,
	}

	fmt.Printf("Length %d\n", rect.Area())

	fmt.Printf("Length %d\n", v.Length)
}
