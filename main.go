package main

import (
	"fmt"
	"strconv"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Fantastic Conference"
var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // NB we changed from a slice of strings to a slice of maps

func main() {

	greetUser()

	for {
		// get user details

		firstName, lastName, email, userTickets := getUserDetails()

		// Validate user details

		detailsCheck := UserDetailsValidation(firstName, lastName, email, userTickets, remainingTickets)
		if !detailsCheck {
			continue
		}

		// Do the transaction:

		updateBookings(userTickets, firstName, lastName, email)

		fmt.Printf("Remaining tickets: %v \n", remainingTickets)

		// firstNames := getFirstNames()
		fmt.Printf("List of bookings clients are: %v \n", bookings)

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
		firstNames = append(firstNames, booking["firstName"])
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

func updateBookings(userTickets uint, firstName string, lastName string, email string) {
	now := time.Now()

	remainingTickets -= userTickets

	// Create a map to store users

	var userData = make(map[string]string)

	// Add user data to map

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is : %v\n", bookings)

	fmt.Printf("Thank you for the purchase %v %v.\nYou bought %v tickets.\nYou will receive a confirmation email to:\n%v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Purchased at: %v:%v on the %v %v %v \n", now.Hour(), now.Minute(), now.Day(), now.Month(), now.Year())

}
