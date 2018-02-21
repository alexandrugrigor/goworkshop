package test

// Rectangle - the rectangle struct
type Rectangle struct {
	Width  int
	Length int
}

// Area - Returns the area
func (r *Rectangle) Area() int {
	return r.Length * r.Width
}

type TestStruct struct{}

// func (r *TestStruct) Area() int {
// 	return 10
// }

// Square - the square struct
type Square struct {
	Rectangle Rectangle
	Color     string
}
