package test

//Rectangle - the rectangle struct
type Rectangle struct {
	Width  int
	Length int
}

//Square - the square struct
type Square struct {
	Rectangle
	Color string
}

//CreateSquare - creates a Square
func CreateSquare(length int) Square {
	if length >= 0 && length <= 10 {
		return Square{
			Rectangle: Rectangle{
				Length: length,
				Width:  length,
			},
			Color: "red",
		}
	}

	return Square{
		Rectangle: Rectangle{
			Length: length,
			Width:  length,
		},
		Color: "Green",
	}
}

//Area - returns the area
func (r *Rectangle) Area() int {
	return r.Length * r.Width
}
