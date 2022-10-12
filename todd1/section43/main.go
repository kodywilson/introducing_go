package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

// sort exercise
type people []string

type Square struct {
	side float64
}

// z Square is the receiver, area() method, returns float64
func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Worker interface {
	Work()
}

type Person struct {
	name string
}

// p Person is the receiver - attach method to Person struct
func (p Person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {
	s := Square{side: 10}
	info(s)
	c := Circle{radius: 10}
	info(c)

	p := Person{
		name: "Bob",
	}
	var w Worker = p
	describe(w)
	w.Work()

	// rest call
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Printf("str is type %T\n", str)
	fmt.Println(str[:100]) // print out first 100 chars

	// sort exercise
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

	numbers := []int{10, 2, 5, 99, 232, 55, 3, 1, 100}
	// sort ints in ascending order
	var bumbers []int
	bumbers = numbers
	sort.Ints(bumbers)
	fmt.Println(numbers) // numbers was sorted, bumbers points to numbers

	// now sort the slice in reverse and then print
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println(numbers)
}
