package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("Type of a is %T and type of b is %T\n", a, b)

	// Use * to read val from address (dereference)
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

}
