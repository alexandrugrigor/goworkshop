package test

//Rectangle
type Rectangle struct {
	Width  int
	Length int
}

func (r *Rectangle) Area() int {
	return r.Length * r.Length
}

type Test struct{}

func (r *Test) Area() int {
	return 2
}

// Square -
type Square struct {
	Rectangle
	Test
	Color string
}
