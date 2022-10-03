package main

import (
	"fmt"
)

// Ex 3
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Exercise 1
	//x := 42
	//y := "James Bond"
	//z := true
	fmt.Println("x, y, z are:", x, y, z)
	fmt.Printf("x is %v of type %T\n", x, x)
	fmt.Println(y)
	fmt.Println(z)

	// Ex 3
	xyz := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(xyz)

	// Exercise 2
	var xx int
	var yy string
	var zz bool
	fmt.Println(xx, yy, zz)
	// These are 0 values (default)

	// Exercise 3
}
