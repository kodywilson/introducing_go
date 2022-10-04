package main

import "fmt"

func main() {
	// boolean type - true / false
	var booly bool
	fmt.Println("booly is", booly)

	a := 7
	b := 42
	fmt.Println(a == b-35)

	// numbers - integers and floats
	x := 42
	y := 42.34534
	fmt.Printf("x type is %T while y type is %T\n", x, y)
}
