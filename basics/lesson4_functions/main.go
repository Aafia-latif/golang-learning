package main

import "fmt"

// Function: get user input
func getUserInput() (string, int) {
	var name string
	var age int

	// Ask for user input
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	return name, age
}

// Function: validate age
func validateAge(age int) bool {
	if age <= 0 {
		fmt.Print("Invalid age entered.")
		return false
	}
	return true
}

// Funtion: determine ticket
func getTicketDetails(age int) (string, int) {
	if age >= 18 {
		return "Adult", 1000
	} else if age >= 13 {
		return "Teen", 700
	}
	return "Child", 500
}

// Function: print ticket details
func printTicket(name string, age int, ticketType string, price int) {
	fmt.Println("\n----- Ticket Details -----")
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Ticket Type: %s\n", ticketType)
	fmt.Printf("Price: PKR %d\n", price)
	fmt.Println("--------------------------")
}

// Function: main
func main() {
	name, age := getUserInput()

	//Validation
	if !validateAge(age) {
		return
	}

	ticketType, price := getTicketDetails(age)
	printTicket(name, age, ticketType, price)
}
