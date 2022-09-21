package main

import "fmt"

func main() {
	var a int = 42
	// use * to create pointer, set type to same as target variable
	var b *int = &a    // use & (address of operator) to get memory address
	fmt.Println(&a, b) // show memory addresses
	// now use dereferencing operator to get value at pointer
	fmt.Println(a, *b) // show memory addresses
	a = 27
	fmt.Println(a, *b)
}
