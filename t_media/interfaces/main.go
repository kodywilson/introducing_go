package main

import (
	"fmt"
	"math"
)

// Define interface
type Shape interface {
	area() float64
}

// shape structs
type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// struct methods
// struct method is on then name of method (area in this case) then type returned
func (c Circle) area() float64 {
	// pi * r * r
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

// now interface method
// I found this very confusing at first, but basically, you are
// creating a function, getArea, that you pass a struct, Shape in this case
// so a circle or rectangle in this example
// the right area() function will be called, depending on which
// shape is passed. Each possible shape needs to have
// an area method defined, you need the interface, and the
// function (method) on the interface
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", getArea(circle))
	fmt.Printf("Rectangle area: %f\n", getArea(rectangle))
}
