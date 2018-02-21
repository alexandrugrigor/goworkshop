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

		Color: "Red",
	}

	fmt.Println(s.Rectangle.Area())
}
