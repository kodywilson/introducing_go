package main

import "fmt"

type Car struct {
	TopSpeed   int
	Name       string
	Cool       bool
	Passengers []string
}

func main() {
	c := Car{
		TopSpeed:   60,
		Name:       "Mirthmobile",
		Cool:       true,
		Passengers: []string{"garth", "wayne"},
	}

	fmt.Println(c)
}
