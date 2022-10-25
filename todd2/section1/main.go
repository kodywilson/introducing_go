package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licenseToKill bool
}

// keyword (receiver) name(arguments) (return) {func body}
func (p person) speak() {
	fmt.Println(p.first, "says Hello")
}

func (sa secretAgent) speak() {
	fmt.Println(sa.first, "says, I am not an agent.")
}

// both persons and secretAgents implement the human interface
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

// exercise 1
type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type shape interface {
	area() float64
}

// function uses the interface
func calcArea(s shape) {
	fmt.Println("Area of shape is", s.area())
}

func main() {
	x := 7
	fmt.Printf("Type of x is %T\n", x)

	//composite
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	// map
	m := map[string]int{
		"Bob":     42,
		"Charles": 65,
	}
	fmt.Println(m)

	// struct
	p1 := person{
		first: "Bob",
		last:  "Johnson",
		age:   25,
	}
	fmt.Println(p1)
	p1.speak()

	sa1 := secretAgent{
		person{
			"James",
			"Bond",
			32,
		},
		true,
	}
	sa1.speak()

	// use human interface
	saySomething(p1)
	saySomething(sa1)

	// ex 1
	c := circle{10}
	calcArea(c)

	s := square{10}
	calcArea(s)

	// ex 2
	// slice of ints
	si := []int{1, 2, 3, 4, 5}
	for i := range si {
		fmt.Println(i)
	}
	for i, v := range si {
		fmt.Println("Index:", i, "Value:", v)
	}

	// map
	mappy := map[string]int{
		"jones":  32,
		"davey":  29,
		"george": 36,
	}
	for k := range mappy {
		fmt.Println(k)
	}
	for k, v := range mappy {
		fmt.Println("Key:", k, "Val:", v)
	}
}
