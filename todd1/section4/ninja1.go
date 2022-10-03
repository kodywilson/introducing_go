package main

import (
	"fmt"
)

// Ex 3
var x int = 42
var y string = "James Bond"
var z bool = true

// Ex 4
type number int

var ex number

// Ex 5
var why int

func main() {
	// Exercise 1
	//x := 42
	//y := "James Bond"
	//z := true
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

	// Ex 3
	xyz := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(xyz)

	// Ex 4
	fmt.Println("ex is", ex)
	fmt.Printf("the type of ex is %T\n", ex)
	ex = 42
	fmt.Println("Now ex is", ex)

	// Ex 5
	why = int(ex)
	fmt.Println("why is", why)
	fmt.Printf("why is type %T\n", why)
}
