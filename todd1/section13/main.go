package main

import "fmt"

// ex1
type person struct {
	firstName     string
	lastName      string
	icecreamFaves []string
}

// ex3
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
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
		bob.firstName:  bob,
		gary.firstName: gary,
	}
	for k, v := range people {
		fmt.Printf("Username: %v\nData:\n", k)
		fmt.Println(v)
	}

	// ex3 - composite literals, embedded structs
	greenie := truck{
		vehicle: vehicle{
			doors: 2,
			color: "green",
		},
		fourWheel: false,
	}

	redracer := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}
	fmt.Println(greenie)
	fmt.Println(redracer)
	fmt.Println(greenie.doors, greenie.fourWheel, redracer.color, redracer.luxury)

	// ex4 - anonymous struct
	anon := struct {
		mappy  map[string]int
		slicey []string
	}{
		mappy: map[string]int{
			"age": 34,
		},
		slicey: []string{"beer", "cheese"},
	}
	fmt.Println(anon)
	fmt.Println(anon.slicey)
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
