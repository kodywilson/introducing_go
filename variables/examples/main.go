package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * input

	fmt.Println(output)
	fmt.Println(a + b + c)
}
