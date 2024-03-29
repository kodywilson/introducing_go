package main

import "fmt"

func main() {
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	// different ways to create arrays
	grades := [3]int{97, 85, 93}
	grades2 := [...]int{87, 95, 76, 99}
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades2: %v\n", grades2)
	students[0] = "Bob"
	students[1] = "Lisa"
	students[2] = "Gerald"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of students: %v\n", len(students))
	a := [...]int{1, 2, 3}
	b := &a
	b[1] = 7
	fmt.Println(a)
	fmt.Println(b)
	// create slice with empty [] or by taking part of an existing array
	c := []int{1, 2, 3}
	fmt.Println(c)
	fmt.Printf("Length: %v\n", len(c))
	fmt.Printf("Capacity: %v\n", cap(c))
	// using make to create array
	d := make([]int, 3, 100)
	fmt.Println(d)
	fmt.Printf("Length: %v\n", len(d))
	fmt.Printf("Capacity: %v\n", cap(d))
	d = append(d, 1)
	fmt.Println(d)
	fmt.Printf("Length: %v\n", len(d))
	fmt.Printf("Capacity: %v\n", cap(d))
	// pop and push
	e := []int{1, 2, 3, 4, 5}
	fmt.Println(e)
	e = append(e, 6) // push an element onto slice
	fmt.Println(e)
	e = e[1:] // pop off the first element
	fmt.Println(e)
	e = e[:len(e)-1] // remove last element
	fmt.Println(e)
	e = []int{1, 2, 3, 4, 5}
	e = append(e[:2], e[3:]...) // remove middle element
	fmt.Println(e)
}
