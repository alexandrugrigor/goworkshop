package test

import (
	"fmt"
	"testing"
)

// 1. How to declare?
type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("Point{x=%d, y=%d}", p.X, p.Y)
}

func (p Point) DoSomething() string {
	return fmt.Sprintf("Point{x=%d, y=%d}", p.X, p.Y)
}

type MyInterface interface {
	DoSomething() string
}

func PrintData(obj MyInterface) {
	fmt.Println(obj.DoSomething())
}
func TestStructs(t *testing.T) {
	// 2. How to initialize?
	p1 := Point{
		X: 10,
		Y: 10,
	}
	PrintData(p1)
	fmt.Printf("\n0. p1 = %v\n", p1)

	p2 := &Point{3, 4}

	fmt.Printf("\n1. p2 = %v\n", p2)

	// p2 initialization is equivalent with p3
	p3 := new(Point)

	fmt.Printf("\n2. p3 = %v\n", p3)
	p3.X = 5
	p3.Y = 6
	// *p3 = Point{5, 6}
	fmt.Printf("\n3. p3 = %v\n", p3)

	// 3. How to embedd struct types?
	// Circle
	type Circle struct {
		Point
		Radius int
	}
	// Wheel
	type Wheel struct {
		Circle
		Spokes int
	}
	wheel := Wheel{
		Circle: Circle{
			Point: Point{
				X: 1,
				Y: 2,
			},
			Radius: 3,
		},
		Spokes: 4,
	}
	fmt.Printf("3. wheel = %v\n", wheel)
	// or equivalent to
	wheel.Point = Point{
		X: 5,
		Y: 6,
	}
	wheel.Radius = 7
	wheel.Spokes = 8
	fmt.Printf("3. wheel = %v\n", wheel)

}
