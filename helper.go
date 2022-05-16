package main

import (
	"fmt"
	"strings"
)

func UserDetailsValidation(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) bool {

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
