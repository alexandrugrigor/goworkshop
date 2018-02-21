package main

import (
	"fmt"
	"goworkshop/test"
)

func main() {
	var s = test.Square{
		Rectangle: test.Rectangle{
			Width:  10,
			Length: 10,
		},
	}
	fmt.Println(s.Test.Area())
}
