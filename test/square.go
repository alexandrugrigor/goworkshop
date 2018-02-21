package test

// Square type
type Square struct {
	Length int
	Color  string
}

type Rectangle struct {
	Square
	Long int
}

// returns the area of the Square
func (s Square) Area() int {
	return s.Length * s.Length
}
