package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	slicey1 := []int{1, 2, 3}
	slicey2 := make([]int, 2)
	copy(slicey2, slicey1)
	fmt.Println(slicey1, slicey2)
}
