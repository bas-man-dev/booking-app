package main

import (
	"fmt"
	"strings"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Fantastic Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUser()

	for {
		// get user details

		firstName, lastName, email, userTickets := getUserDetails()

		// Validate user details

		detailsCheck := userDetailsValidation(firstName, lastName, email, userTickets)
		if !detailsCheck {
			continue
		}

		// Do the transaction:

		updateBookings(userTickets, firstName, lastName, email)

		fmt.Printf("Remaining tickets: %v \n", remainingTickets)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings clients are: %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Printf("%v fully booked... See you again next year!\n", conferenceName)
			break
		}

	}

}

func greetUser() {

	fmt.Printf("\nWelcome to the %v booking application:\n", conferenceName)
	fmt.Printf("We have %v tickets and %v available\n", conferenceTickets, remainingTickets)
	fmt.Println("Continue here to get your tickets to attend the conference")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserDetails() (string, string, string, uint) {
	firstName := ""
	lastName := ""
	email := ""
	var userTickets uint

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address: ")
	fmt.Scan(&email)

	fmt.Printf("%v Tickets left. How many tickets would you like? \n", remainingTickets)
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func userDetailsValidation(firstName string, lastName string, email string, userTickets uint) bool {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTickets

	if !isValidName {
		fmt.Println("First name and Last name need to be 2 or more characters long")
		return false
	}
	if !isValidEmail {
		fmt.Println("email address needs an '@' sign")
		return false
	}
	if !isValidTickets {
		fmt.Println("Ticket number was invalid.")
		return false
	}
	return true
}

func updateBookings(userTickets uint, firstName string, lastName string, email string) {
	now := time.Now()

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you for the purchase %v %v.\nYou bought %v tickets.\nYou will receive a confirmation email to:\n%v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Purchased at: %v:%v on the %v %v %v \n", now.Hour(), now.Minute(), now.Day(), now.Month(), now.Year())

}
