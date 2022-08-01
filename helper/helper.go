package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, ticketsRemaining uint) (bool, bool, bool) {
	isValName := len(firstName) >= 2 && len(lastName) >= 2
	isValEmail := strings.Contains(email, "@")
	isValTicketNumber := userTickets > 0 && userTickets <= ticketsRemaining
	return isValName, isValEmail, isValTicketNumber
}
