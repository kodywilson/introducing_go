package main

import "fmt"

func main() {
	var x [10]int
	x[3] = 30
	fmt.Println(x[3])

	slice1 := []int{1, 2, 3}
	fmt.Println(slice1[0])

	slice2 := make([]int, 3, 9)
	fmt.Println(len(slice2))
}
