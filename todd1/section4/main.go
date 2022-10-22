package main

import "fmt"

type hotdog int

var a = 42
var b hotdog = 420

func main() {
	x := `"This is a quote"
	and it appears in a string literal`
	fmt.Println(x)
	fmt.Printf("x is a %T\n", x)
	y := 42
	fmt.Printf("x in base 10: %d and in binary: %b and in hex: %#X\n", y, y, y)
	z := fmt.Sprintf("%d = %b = %#X", y, y, y)
	fmt.Println(z)
	fmt.Printf("variable b is type %T\n", b)
	a = int(b)
	fmt.Printf("variable a is type %T\n", a)
}
