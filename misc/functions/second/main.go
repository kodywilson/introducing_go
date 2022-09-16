package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f(x int) {
	fmt.Println(x)
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}

func f3() (int, int) {
	return 5, 6
}

func main() {
	xs := []float64{98, 93, 77, 82, 83, 1000, 24, 100000}
	fmt.Println(average(xs))
	x := 5
	f(x)
	fmt.Println(f1())
	x, y := f3()
	fmt.Println(x, y)
}
