package main

import "fmt"

var (
	a = 5
	b = 10
	c = 15
)

func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * input

	fmt.Println(output)
	fmt.Println(a + b + c)
	fmt.Println(a * b * c)

	x := 5
	x += 1
	fmt.Println(x)

	fmt.Print("Enter a temperature in Fahrenheit: ")
	var f_temp float64
	fmt.Scanf("%f", &f_temp)

	fmt.Println("In Celcius, that would be: ")
	fmt.Println(f2c(f_temp))

	fmt.Print("Enter a length in feet: ")
	var f_length float64
	fmt.Scanf("%f", &f_length)

	fmt.Println("In meteres, that would be: ")
	fmt.Println(feet2meters(f_length))
}

func f2c(f float64) float64 {
	celc := (f - 32) * 5 / 9
	return celc
}

func feet2meters(f float64) float64 {
	metey := f * 0.3048
	return metey
}
