package main

import "goworkshop/test"
import "fmt"

func main() {
	fmt.Println("hello world!")
	fmt.Println("Catalin")
	fmt.Println("Hi, my name is Alex")
	fmt.Println("Hi, my name is Tibi")
	fmt.Println("Hi, my name is Anda! :)")

	// var s = test.createSquare(11)
	// s.Color = "green"
	// fmt.Println(s)

	// var s = test.Square{
	// 	Length: 10,
	// }
	// s.Color = "green"

	var s = test.CreateSquare(10)
	fmt.Println(s.Area())

	var book = BookDto{
		Title: "A Clockwork Orange",
	}
	fmt.Println(book)

}
