package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	sum := 0
	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
		sum += id
	}
	fmt.Println("Sum:", sum)
}
