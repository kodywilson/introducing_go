package main

import "fmt"

func main() {
	// array
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x)) // use len() to get the length
	// slices are more useful. They are backed by arrays
	// use composite literal syntax: type {}
	sl := []int{} // you don't have to specify length or even starting values
	for i := 0; i < 8; i++ {
		fmt.Printf("Length of slice is %v\n", len(sl))
		sl = append(sl, i) // slice grows dynamically
	}
	// one way to loop over slice
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
	// another way to loop over slice
	for i, v := range sl { // i = index, v = value
		fmt.Printf("Index %v: %v\n", i, v)
	}
	// access one value at specific index
	fmt.Println(sl[7])
}
