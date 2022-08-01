package main

import (
	"booking-tickets/helper"
	"fmt"
	"strings"
)

const confTickets int = 50

var confName = "Go Conference"

var ticketsRemaining uint = 50
var bookings []string

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValName, isValEmail, isValTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets)

		if isValName && isValEmail && isValTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			//call function print firstnames
			firstName := getFirstNames()
			fmt.Printf("The first names of booking are: %v\n", firstName)
			if ticketsRemaining == 0 {
				fmt.Println("Our conference is now booked out. Please come back next year.")
				break
			}

		} else {
			if !isValName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValTicketNumber {
				fmt.Println("number of tickets you entered is invailed")
			}
		}
	}
}
func greetUsers() {
	fmt.Printf("Welcome to our %v booking application. You have %v tickets available and %v remaining\n", confName, confTickets, ticketsRemaining)

}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var email string
	var firstName string
	var lastName string
	var userTickets uint
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Please enter your no. of tickets:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	ticketsRemaining += -userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n ", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", ticketsRemaining, confName)
}
