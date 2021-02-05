package main

import (
	"fmt";
	"math"
)

// square is a square
type Square struct {
	Length float64
}

// area returns the area of the square
func (s *Square) Area() float64 {
	return s.Length * s.Length
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Return area of circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// shape requied!!!
func sumAreas (shapes []Shape) float64 {
	total := 0.0

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}

// INTERFACE: shape
type Shape interface {
	Area() float64
}

func main() {
	s := &Square{20}
	fmt.Println(s.Area())

	c := &Circle{10}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sa := sumAreas(shapes)
	fmt.Println(sa)
}
