package main

import "fmt"

func main() {
	ids := []int{33, 76, 54, 23, 11, 2}

	sum := 0
	// Loop through slice of ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
		sum += id
	}
	fmt.Println("Sum:", sum)

	// Loop through map
	books := map[string]int{
		"Chamber of secrets": 256,
		"Lord of the Rings":  735,
	}

	for k, v := range books {
		fmt.Println(k, v)
	}
}
