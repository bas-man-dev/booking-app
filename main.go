package main

import (
	"fmt"
	"time"
)

const conferenceTickets uint = 50

var conferenceName = "Fantastic Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // NB we changed from a slice of maps to a slice of struct {UserData}
// create a struct for different variable types

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

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
		// fmt.Printf("List of bookings clients are: %v \n", bookings)

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
		firstNames = append(firstNames, booking.firstName)
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

	// Create a map to store users *Updated to struct*

	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	// Removed the previous key:value inits

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is : %v\n", bookings)

	fmt.Printf("Thank you for the purchase %v %v.\nYou bought %v tickets.\nYou will receive a confirmation email to:\n%v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Purchased at: %v:%v on the %v %v %v \n", now.Hour(), now.Minute(), now.Day(), now.Month(), now.Year())

}
