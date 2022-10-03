package main

import "fmt"

func main() {
	x := `"This is a quote"
	and it appears in a string literal`
	fmt.Println(x)
	fmt.Printf("x is a %T\n", x)
}
