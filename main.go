package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets int = 50
var bookings = []string{}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidTicketsNumber, isValidEmail := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidTicketsNumber && isValidEmail {
			bookTicket(userTickets, firstName, lastName, conferenceName, email)

			var firstNames = getFirstNames()
			fmt.Printf("The first names of bookings are : %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name is invalid")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid")
			}
			if !isValidTicketsNumber {
				fmt.Println("TicketsNumber is invalid")
			}

		}
	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v of tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketsNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketsNumber
}
func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, conferenceName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
