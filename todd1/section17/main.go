package main

import "fmt"

// ex2
type person struct {
	first string
}

func main() {
	// ex1
	x := 3
	fmt.Println(&x)

	// ex2
	p := person{first: "Tommy"}
	fmt.Println(p.first)
	changeMe(&p)
	fmt.Println(p.first)
}

// ex2
func changeMe(p *person) {
	p.first = "Bobbers"
	// could also do (*p).first = "Bob"
}
