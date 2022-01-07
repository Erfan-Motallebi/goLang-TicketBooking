package main

import (
	"fmt"
	"strings"
)

func main() {

	var ticketNumbers uint = 50
	var remainingTickets uint = 0
	fullNames := []string{}

	cities := []string{"London", "USA", "Iran", "Canada", "Germany", "Japan", "Korea"}

	fmt.Println("Please pick a city from the list", cities)

	for {

		// City Checker
		var isCityOk bool = cityCheck(cities)

		if !isCityOk {
			fmt.Println("Please Pick an available city for taking a room please.")
			continue
		}
		// User Inputs
		firstName, lastName, email, rooms := userInput()
		fmt.Printf("FirstName: %s - lastName: %s - Email: %v\n", firstName, lastName, email)

		// User Welcomer
		userWelcomer(firstName)
		// User Input Validations

		isValidFirstname, isValidLastname, isValidEmail, isValidState := userValidator(firstName, lastName, email, ticketNumbers, remainingTickets, rooms)
		// Checking the state of user inputs
		switch {
		case !isValidFirstname:
			{
				fmt.Println("Wrong surname")
				continue
			}
		case !isValidLastname:
			{
				fmt.Println("Wrong lastname")
				continue
			}
		case !isValidEmail:
			{
				fmt.Println("Wrong Email")
				continue
			}
		case !isValidState:
			{
				fmt.Println("Bad state")
				continue
			}
		default:
			{
				fmt.Println("Successfully went through the precedure")
			}
		}

		fullNames = append(fullNames, firstName+" "+lastName)

		// Extracting the firstname of the customers
		firstNames := firstNameExtractor(fullNames)
		// figuring out the ticket numbers and showing the final state
		ticketNumbers = ticketNumbers - rooms
		fmt.Println("the first name of customers would be like: ", firstNames)
		fmt.Printf("Available tickets : %v\n", ticketNumbers)
		continue
	}

}

func userInput() (string, string, string, uint) {
	var email string
	var firstName string
	var lastName string
	var rooms uint
	fmt.Println("Please enter your name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your lastname:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Please enter your rooms: ")
	fmt.Scan(&rooms)

	return firstName, lastName, email, rooms
}

func userValidator(firstName, lastName, email string, ticketNumbers, remainingTickets, rooms uint) (bool, bool, bool, bool) {
	var isValidFirstname bool = len(firstName) > 2
	isValidLastname := len(lastName) > 3
	isValidState := ticketNumbers < remainingTickets
	isValidEmail := strings.Contains(email, "@")
	remainingTickets = ticketNumbers - rooms

	return isValidFirstname, isValidLastname, isValidEmail, isValidState
}

func firstNameExtractor(fullNames []string) []string {
	var firstNames []string
	for _, fullnames := range fullNames {
		var fname = strings.Split(fullnames, " ")
		firstNames = append(firstNames, fname[0])
	}
	return firstNames
}

func userWelcomer(name string) {
	fmt.Printf("Welcome to your conference: Mr.%s\n", name)
	fmt.Println("I wish you all the best of two worlds")
}

func cityCheck(cities []string) bool {
	var item string
	fmt.Scan(&item)

	for i := 0; i < len(cities); i++ {
		if cities[i] == item {
			return true
		}
	}
	return false

}
