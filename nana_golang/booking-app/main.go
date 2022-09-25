package main

import (
	"fmt"
	"strings"
  "booking-app/helper" // module name + directory
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var	remainingTickets = 50
// slice to hold bookings
var bookings []string

func main() {
	
	greetUsers()

	for {
		var exit_app string

		firstName, lastName, email, userTickets := getUserInput()

		bookTicket(userTickets, firstName, lastName, email)

		// Leave app if tickets are all gone
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

func greetUsers() {
	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) // fields will split on space
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their first name, validate input
	for {
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
		// package.ExportedFuncName - when we moved validateName to helper package, we had to capitalize
		// to export it for use in another package and we call it with package.FuncName
		if !helper.ValidateName(firstName) {
			fmt.Printf("You entered %v, first name must be at least 2 and no more than 32 characters\n", firstName)
		} else {
			break
		}
	}

	//ask user for their last name, validate input
	for !helper.ValidateName(lastName) {
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)
		if !helper.ValidateName(lastName) {
			fmt.Printf("You entered %v, last name must be at least 2 and no more than 32 characters\n", lastName)
		}
	}

	//ask for email, validate input
	for !strings.Contains(email, "@") || !strings.Contains(email, ".") || len(email) < 5 {
		fmt.Println("Enter your email address:")
		fmt.Scan(&email)
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") || len(email) < 5 {
			fmt.Printf("You entered %v, email must contain @ and . in address\n", email)
		}
	}

	// ask for number of tickets, validate input
	for {
		fmt.Println("Enter number of tickets you would like to buy:")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
		} else if userTickets <= 0 {
			fmt.Printf("Please enter a non zero, non negative number up to remaining tickets: %v\n", remainingTickets)
		} else {
			break
		}
	}

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets int, firstName string, lastName string, email string) {
	// Confirmation message
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("First names of those attending: %v\n", getFirstNames())

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
