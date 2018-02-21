package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	fmt.Println("Mircea")

	var s = Square {
		Length: 10,
	}
	fmt.Println(s)

}

// Square - the square struct
type Square struct {
	Length int 
	Color string
}

func CreateSquare(length int) Square{
	if length >= 0 && length <= 10 {
		return Square{
			Length: length,
			Color: "red",
		}
	}
	return Square
}