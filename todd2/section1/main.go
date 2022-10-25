package main

import "fmt"

func main() {
	x := 7
	fmt.Printf("Type of x is %T\n", x)

	//composite
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	// map
	m := map[string]int{
		"Bob":     42,
		"Charles": 65,
	}
	fmt.Println(m)
}
