package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("For the variable a, the value %v is stored at %v\n", a, &a)
	fmt.Printf("Type of a is %T while type of &a is %T\n", a, &a)
	b := &a // b points to the memory address of a
	fmt.Println(b)
	fmt.Println(*b)
	*b = 33        // * dereferences when used this way
	fmt.Println(a) // we changed the value of a
	// passing pointers can save lots of memory

	// Go passes by value
	x := 0
	foo(x)
	fmt.Println(x) // prints zero again

	y := 0
	fuu(&y)
	fmt.Println(y)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

// takes a pointer to int
func fuu(y *int) {
	fmt.Println(*y)
	*y = 43
	fmt.Println(*y)
}
