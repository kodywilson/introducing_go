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

}
