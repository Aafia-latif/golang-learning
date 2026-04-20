package main

import "fmt"

// Struct (data model)
type User struct {
	name       string
	age        int
	ticketType string
	price      int
}

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

// Function: assign ticket details
func assignTicket(user *User) {
	if user.age >= 18 {
		user.ticketType = "Adult"
		user.price = 1000
	} else if user.age >= 13 {
		user.ticketType = "Teen"
		user.price = 700
	} else {
		user.ticketType = "Child"
		user.price = 500
	}
}

// Function: print ticket
func printTicket(user *User) {
	fmt.Println("\n----- Ticket Details -----")
	fmt.Printf("Name: %s\n", user.name)
	fmt.Printf("Age: %d\n", user.age)
	fmt.Printf("Ticket Type: %s\n", user.ticketType)
	fmt.Printf("Price: PKR %d\n", user.price)
	fmt.Println("--------------------------")
}

func main() {
	var user User

	// Input
	name, age := getUserInput()

	// Validation
	if !validateAge(age) {
		fmt.Println("Invalid age.")
		return
	}

	// Assign values
	user.name = name
	user.age = age

	// Business logic
	assignTicket(&user)

	// Output
	printTicket(&user)
}
