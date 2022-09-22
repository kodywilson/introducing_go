package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	GrantType   string `yaml:"grant_type"`
	ContentType string `yaml:"content_type"`
	//Authorization string `yaml:"authorization"` // use env variable while testing
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty Config to be are target of unmarshalling
	var c Config

	// Unmarshal our input YAML file into empty Config (var c)
	if err := yaml.Unmarshal(f, &c); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	fmt.Printf("%+v\n", c)
	// Print out one member of struct
	//fmt.Printf("%v\n", c.Authorization)
}
