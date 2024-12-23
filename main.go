package main

import (
	"fmt"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets int = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	number0Tickets int
}

func main() {
	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidTicketsNumber, isValidEmail := ValidateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidTicketsNumber && isValidEmail {
			bookTicket(userTickets, firstName, lastName, conferenceName, email)
			go sendTicket(userTickets, firstName, lastName, email)
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

// Функция приветствие
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total %v of tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Функция ввода данных от User
func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

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

// Добавляем Имю и Фамилия Usera, Благодарность за покупку билета + количество оставшихся билетов
func bookTicket(userTickets int, firstName string, lastName string, conferenceName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Добавляем карту пользователя

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		number0Tickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Отделяем Имя от фамилии для вывода имени
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("####################")
}
