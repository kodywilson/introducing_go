package main

import "fmt"

func main() {
	// Array
	var fruitArr [2]string
	veggieArr := [2]string{"Carrot", "Lettuce"}

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)
	fmt.Println(veggieArr)

	// Slices
	fruitSlice := []string{"Apple", "Banana", "Orange"}

	fmt.Println(fruitSlice)
	fmt.Printf("fruitSlices is type %T\n", fruitSlice)

	// Get part of a slice
	fmt.Println(fruitSlice[:2])
}