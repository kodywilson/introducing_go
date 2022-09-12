package main

import "fmt"

func main() {
	var x string = "Hello, World"
	fmt.Println(x)
	var y string
	y = "Nice to meet you"
	fmt.Println(y)
	var z string
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)
	z = z + " third"
	fmt.Println(z)
	z += " fourth"
	fmt.Println(z)
	var a string = "hello"
	var b string = "world"
	fmt.Println(a == b)
	b = "hello"
	fmt.Println(a == b)
}
