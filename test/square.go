package test

//Rectangle- the rectangle struct
type Rectangle struct {
	Width  int
	Length int
}

//Area- returns the area
func (r *Rectangle) Area() int {
	return r.Length * r.Width
}

//Square- the square struct
type Square struct {
	Rectangle Rectangle
	Color     string
}

type TestStruct struct {
	// func (r *TestStruct) Area() int{
	// 	return 10
	// }
}
