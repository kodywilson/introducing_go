package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

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

	out, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
