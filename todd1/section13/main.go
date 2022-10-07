package main

import "fmt"

type person struct {
	firstName     string
	lastName      string
	icecreamFaves []string
}

func main() {
	// ex 1 - create struct with slice of strings
	bob := person{
		firstName:     "Bob",
		lastName:      "Johnson",
		icecreamFaves: []string{"rocky road", "chocolate mint"},
	}
	gary := person{
		firstName:     "Gary",
		lastName:      "Johnson",
		icecreamFaves: []string{"vanilla", "beach day"},
	}
	//fmt.Println(bob, gary)
	favoriteIcecream(bob)
	favoriteIcecream(gary)

	// ex2 - store structs in map
	people := map[string]person{
		"Bob":  bob,
		"Gary": gary,
	}
	for k, v := range people {
		fmt.Printf("Username: %v\nData:\n", k)
		fmt.Println(v)
	}
}

// ex 1
func favoriteIcecream(p person) {
	fmt.Printf("%v's favorite icecream flavors are: ", p.firstName)
	for i, v := range p.icecreamFaves {
		fmt.Printf("%v", v)
		if i < len(p.icecreamFaves)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
