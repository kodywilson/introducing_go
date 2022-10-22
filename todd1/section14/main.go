package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// method signature
// func (r receiver) identfier(parameters) (return(s)) { code }
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "I'm not a secret agent.")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "I am a secret agent.")
}

// interfaces allow us to define behavior and use polymorphism
type human interface {
	speak()
}

// any object that implements the interface methods can be used as that type

func talk(h human) {
	h.speak()
	switch h.(type) { // switch on type
	case person: // assertion lets you get back to specific type
		fmt.Printf("I am a person and my first name is %v\n", h.(person).first) // assertion that this is a person
	case secretAgent:
		fmt.Printf("I am a secret agent and my first name is %v\n", h.(secretAgent).first)
	}
}

// attach to receiver (secretAgent in this case)

func main() {
	defer foo() // defer causes function to run at end of enclosing function
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x, y)
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	boo(xi...) // send variable amount of ints or use ... after slice
	// to unfurl the slice

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	p1 := person{
		first: "Dr",
		last:  "No",
	}

	fmt.Println(p1)
	talk(p1)
	talk(sa1)

	// anonymous functions
	func() {
		fmt.Println("Anonymous func!")
	}() // need () to run it

	func(x int) {
		fmt.Println("The meaning of life:", x)
	}(42)

	f := func() {
		fmt.Println("this is a func expression")
	}
	f()
	fmt.Printf("%T\n", f)

	// return func()
	ex := fooo()
	fmt.Printf("%T\n", ex)

	i := ex()
	fmt.Println(i)

	// callbacks - pass function as parameter
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9))
	// we pass in the function sum and a list of numbers to even
	fmt.Println("Add up only even numbers:", even(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9))
	// we pass in the function sum and a list of numbers to odd
	fmt.Println("Add up only odd numbers:", odd(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9))

	// closure
	a := incrementor()
	b := incrementor()
	// they each get separate value of x within the called function
	fmt.Println(a()) // 1
	fmt.Println(a()) // 2
	fmt.Println(a()) // 3
	fmt.Println(a()) // 4
	fmt.Println(b()) // 1
	fmt.Println(b()) // 2

	// recursion
	fmt.Println(factorial(4))
	// no recursion - loop
	fmt.Println(loopFactorial(4))
}

// func (r receiver) identifier(parameters) (return(s)) { ... }

func foo() {
	fmt.Println("Hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}

// variadic parameter - the variadic has to be the final parameter!
func boo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}

// this function returns a function that returns an int
func fooo() func() int {
	return func() int {
		return 2000
	}
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}

	return total
}

// take in function as argument that you can then
// use in this function
func even(f func(xi ...int) int, vi ...int) int {
	var evens []int
	for _, v := range vi {
		if v%2 == 0 {
			evens = append(evens, v)
		}
	}
	return f(evens...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var odds []int
	for _, v := range vi {
		if v%2 != 0 {
			odds = append(odds, v)
		}
	}
	return f(odds...)
}

// closure
func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

// recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	total := 1
	for i := 1; i <= n; i++ {
		total *= i
	}
	return total
}
