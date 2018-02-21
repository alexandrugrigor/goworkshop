package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a1 = AuthorDto{
		UUID:      "1",
		FirstName: "Dorinel",
		Lastname:  "Panaite",
		Birthday:  "08.06.1997",
		Death:     "???",
	}

	var a2 = AuthorDto{
		UUID:      "2",
		FirstName: "Dorinelll",
		Lastname:  "Panaiteee",
		Birthday:  "07.06.1997",
		Death:     "???",
	}

	var b1 = BookDto{
		UUID:        "book1",
		Title:       "title1",
		NoPages:     10,
		ReleaseDate: "2017",
		Author:      a1,
	}

	var b2 = BookDto{
		UUID:        "book2",
		Title:       "title2",
		NoPages:     20,
		ReleaseDate: "2016",
		Author:      a2,
	}

	authors := [2]AuthorDto{a1, a2}
	books := [2]BookDto{b1, b2}

	b1json, _ := json.Marshal(b1)

	fmt.Println(authors)
	fmt.Println(books)
	fmt.Println(b1json)

	var b3 = new(BookDto)
	error := json.Unmarshal(b1json, &b3)

	fmt.Println(error)
	fmt.Println(b3)
}
