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
	// slice a slice
	fmt.Println(sl[3:7]) // : operator [start:end] (up to but not including)
	fmt.Println(sl[4:])  // from 4 to end
	fmt.Println(sl[:4])  // from start to 4 (up to, but not including)
	// append to a slice
	sl = append(sl, 8, 9, 10)
	sl2 := []int{}
	sl2 = append(sl2, sl...) // use ... to unpack the values from the slice
	fmt.Println(sl2, len(sl2))
	// remove from slice
	// you use append and slice the slice for vals in the middle
	sl2 = sl2[:8] //append(sl2[:8], sl2[8:]...)
	fmt.Println(sl2, len(sl2))
	// using make to create slices
	sl3 := make([]int, 10, 10) // type, values, capacity
	fmt.Println(sl3, len(sl3))
	sl3 = append(sl3, 1, 1)    // this used to double the cap if
	fmt.Println(sl3, len(sl3)) // it went past current cap

	// multi dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "tequila"}
	fmt.Println(jb, mp)
	xp := [][]string{jb, mp} // multi-dimensional slice
	fmt.Println(xp)
	fmt.Println(xp[0])    // first slice of slices
	fmt.Println(xp[0][1]) // second slice of slice of slices

	// maps
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	if v, ok := m["Bob"]; ok {
		fmt.Println("Bob is here too and is", v, "years old")
	} else {
		fmt.Println("Bob isn't here yet!")
	}
	// adding to maps
	m["Bob"] = 5
	if v, ok := m["Bob"]; ok {
		fmt.Println("Bob is here now and is", v, "years old")
	}
	// range over values
	for k, v := range m {
		fmt.Printf("Key: %v\tAge: %v\n", k, v)
	}
	fmt.Println("Now we remove Bob...")
	// delete entry from map
	delete(m, "Bob")      // can use , ok with if statement to see if
	for k, v := range m { // value is there before deleting
		fmt.Printf("Key: %v\tAge: %v\n", k, v)
	}
}
