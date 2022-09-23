package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50
	// array to hold bookings
	var bookings []string

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var exit_app string
		var firstName string
		var lastName string
		var email string
		var userTickets int
		//ask user for their name and email
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets you would like to buy:")
		fmt.Scan(&userTickets)

		// Confirmation message
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking) // fields will split on space
			firstNames = append(firstNames, names[0])
		}
		fmt.Println(firstNames)
		fmt.Printf("These are all our bookings: %v\n", bookings)

		if remainingTickets == 0 {
			fmt.Printf("%v is sold out! Come back next year.\n", conferenceName)
			break
		}
		fmt.Println("Would you like to exit (Y/N)?")
		fmt.Scan(&exit_app)
		if exit_app == "Y" || exit_app == "y" {
			break
		}
	}
}
