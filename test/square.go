package test

// type Square struct {
// 	Length int
// 	Color  string
// }

func CreateSquare(length int) Square {
	if length >= 0 && length <= 10 {
		return Square{
			Rectangle: Rectangle{Width: length, Length: length},
			Color:     "red",
		}
	} else {
		return Square{
			Rectangle: Rectangle{Width: length, Length: length},
			Color:     "blue",
		}
	}
}

func (r *Rectangle) Area() int {
	return r.Width * r.Length
}

type Rectangle struct {
	Length int
	Width  int
}

type Square struct {
	Rectangle
	Color string
}
