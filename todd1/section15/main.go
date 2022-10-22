package main

import (
	"fmt"
)

// ex 4 - create person struct and attach method
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "and I am", p.age, "years old.")
}

// ex5 - create circle, square, shape interface and area methods
type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func (s square) area() float64 {
	return s.side * s.side
}

type shape interface {
	area() float64 // this is the method signature
} // to use this interface, match the method signature

func info(s shape) {
	fmt.Println("Area is", s.area())
}

func main() {
	// ex3
	defer deferred() // will not run until the end of main

	// ex1 - two functions, one returns int, other returns int and string
	fmt.Println(foo())
	fmt.Println(bar())

	// ex2
	si := []int{1, 2, 3}
	fmt.Println(fuu(si...))
	fmt.Println(baar(si))

	// ex4 - struct and method on struct
	p1 := person{
		first: "Bob",
		last:  "Johnson",
		age:   23,
	}
	p1.speak()

	// ex5 - structs, interface methods
	c := circle{radius: 10}
	s := square{side: 10}
	info(c)
	info(s)

	// ex6 - anonymous func
	func() {
		fmt.Println("Anonymous func!")
	}()

	// ex7 - assign func to variable and call that func
	f := func(name string) {
		fmt.Println("Hello,", name)
	}
	f("Bob")

	// ex8 - create a function that returns a function
	fun := funky()
	fmt.Println(fun())

	// ex9 - callback - pass a func into a func as an argument
	// I pass sum func into monkey and use it in monkey
	fmt.Println(monkey(sum, 1, 2, 3))

	// ex10 - closure
	a := closey()
	b := closey()
	a()
	a()
	a()
	b()
	b()
	fmt.Println(a())
	fmt.Println(b())
}

// ex1
func foo() int {
	return 42
}

func bar() (int, string) {
	return 7, "Lucky!"
}

// ex2
func fuu(n ...int) int {
	fmt.Printf("%T\n", n) // type is slice
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}

func baar(si []int) int {
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

// ex3 - defer this func
func deferred() {
	defer fmt.Println("Me too!") // will run at end of func deferred
	fmt.Println("I should have gone first, but I got deferred!")
}

// ex8 - return a func
func funky() func() int {
	return func() int {
		return 42
	}
}

// ex9 - callback - pass func into func as an argument
func monkey(f func([]int) int, si ...int) int {
	return f(si)
}

func sum(si []int) int {
	total := 0
	for _, v := range si {
		total += v
	}
	return total
}

// ex10 - closure
func closey() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
