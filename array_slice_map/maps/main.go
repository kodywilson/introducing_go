package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 356
	fmt.Println(y[1])

	z := make(map[int]int)
	for i := 1; i < 11; i++ {
		z[i] = i * 10
		fmt.Println("The length of z is: ", len(z))
		fmt.Println(z[i])
	}

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	if name, ok := elements["Ne"]; ok {
		fmt.Println(name, ok)
	}

}
