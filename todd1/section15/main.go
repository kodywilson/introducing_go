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

	// ex4
	p1 := person{
		first: "Bob",
		last:  "Johnson",
		age:   23,
	}
	p1.speak()
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
