package main

import (
	"fmt"
)

func main() {
	var s = Square{
		Length: 10,
		Color:  "red",
	}

	fmt.Println(s)
}

// Square - the square struct
type Square struct {
	Length int // upper-case public, lower-case package-private
	Color  string
}
