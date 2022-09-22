package main

import "fmt"
 
func main() {
	x := 10
	y := 15

	// if else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "black"

	if color == "red" {
		fmt.Println("color is", color)
	} else if color == "blue" {
		fmt.Println("color is", color)
	} else {
		fmt.Println("color is not red or blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}

}