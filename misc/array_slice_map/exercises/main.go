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

	y := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(y[2:5])
	fmt.Println(y[0:6])

	z := []int{
		48, 96, 86, 68, 3,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	numby := 0
	lowest := false
	for i := 0; i < len(z); i++ {
		for j := 0; j < len(z); j++ {
			fmt.Println(z[i], z[j])
			if z[i] > z[j] {
				lowest = false
				break
			}
			lowest = true
		}
		if lowest == true {
			numby = z[i]
		}
	}
	fmt.Println(numby)
}
