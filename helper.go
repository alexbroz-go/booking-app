package main

import "strings"

// Проверка на валидность вводимых данных
func ValidateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	var isValidName = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	var isValidTicketsNumber = userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketsNumber
}
