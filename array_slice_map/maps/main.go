package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 356
	fmt.Println(y[1])
}
