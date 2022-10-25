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
}
