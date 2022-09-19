package main

import "fmt"

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	// Create map of states with population
	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	fmt.Printf("The population of Texas is %v\n", statePopulations["Texas"])
	statePopulations["Georgia"] = 10310371 // add a new state
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia") // remove state
	fmt.Println(statePopulations)
	_, ok := statePopulations["Ohio"] // check if key is present
	fmt.Println(ok)
	fmt.Println(len(statePopulations))
	// maps pass by reference
	sp := statePopulations
	delete(sp, "Ohio") // change sp and you change statePopulations
	fmt.Println(sp)
	fmt.Println(statePopulations)

	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah",
			"Jane Smith",
		},
	}
	// You can do this, but it's probably not a good idea...
	bDoctor := Doctor{4, "Jack Holiday", []string{"Billy", "Bob"}}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[0])
	fmt.Println(bDoctor)
	// structs pass by copy
	cDoctor := bDoctor
	cDoctor.actorName = "Jazz Handy"
	fmt.Println(cDoctor)
}
