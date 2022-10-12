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

func talk(h human) {
	h.speak()
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
