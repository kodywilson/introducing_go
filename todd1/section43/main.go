package main

import "fmt"

type Square struct {
	side float64
}

// z Square is the receiver, area() method, returns float64
func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	s := Square{side: 10}
	info(s)
	c := Circle{radius: 10}
	info(c)
}
