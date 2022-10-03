package main

import (
	"fmt"
)

func main() {
	// Exercise 1
	x := 42
	y := "James Bond"
	z := true
	fmt.Println("x, y, z are:", x, y, z)
	fmt.Printf("x is %v of type %T\n", x, x)
	fmt.Println(y)
	fmt.Println(z)

	// Exercise 2
	var xx int
	var yy string
	var zz bool
	fmt.Println(xx, yy, zz)
	// These are 0 values (default)
}
