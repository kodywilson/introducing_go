package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 356
	fmt.Println(y[1])

	z := make(map[int]int)
	for i := 0; i < 10; i++ {
		z[i] = i * 10
		fmt.Println("The length of z is: ", len(z))
		fmt.Println(z[i])
	}

}
