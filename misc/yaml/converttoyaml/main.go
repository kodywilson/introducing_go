package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Car struct {
	TopSpeed   int      `yaml:"topspeed"`
	Name       string   `yaml:"name"`
	Cool       bool     `yaml:"cool"`
	Passengers []string `yaml:"passengers"`
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("yammy.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty Car to be are target of unmarshalling
	var c Car

	// Unmarshal our input YAML file into empty Car (var c)
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	fmt.Printf("%+v\n", c)
}
