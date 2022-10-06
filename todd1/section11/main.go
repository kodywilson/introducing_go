package main

import "fmt"

func main() {
	// section 11
	// ex 1 - create array and range over it
	a := [5]int{1, 2, 3, 4, 5}
	for i := range a {
		fmt.Println(a[i])
	}
	fmt.Printf("%T\n", a)

	// ex2 - create slice and range over it
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := range sl {
		fmt.Println(sl[i])
	}
	fmt.Printf("%T\n", sl)

	// ex3 - slice a slice
	fmt.Println(sl[:6])
	fmt.Println(sl[6:])
	fmt.Println(sl[2:4])
	fmt.Println(sl[3:4])

	// ex4 - append to slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 52, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...) // unfurl slice into target slice
	fmt.Println(x)

	// ex5 - delete from slice
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)

	// ex6 - create slice of strings with make
	states := make([]string, 0, 50)
	states = append(states, "Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado")
	fmt.Println(states)
	fmt.Println(len(states))
	for i := 0; i < len(states); i++ {
		fmt.Println(states[i])
	}
}
