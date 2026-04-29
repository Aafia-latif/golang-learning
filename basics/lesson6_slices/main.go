package main

import "fmt"

// Struct: User data model
type User struct {
	name       string
	age        int
	ticketType string
	price      int
}

// Input function
func getUserInput() (string, int) {
	var name string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	return name, age
}

// Validation
func validateAge(age int) bool {
	return age > 0
}

// Assign ticket (business logic)
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

// Print single ticket
func printTicket(user User) {
	fmt.Println("\n----- Ticket -----")
	fmt.Printf("Name: %s | Age: %d\n", user.name, user.age)
	fmt.Printf("Type: %s | Price: PKR %d\n", user.ticketType, user.price)
}

// Print all users
func printAllUsers(users []User) {
	fmt.Println("\n===== All Tickets =====")
	for _, user := range users {
		printTicket(user)
	}
}

// MAIN FUNCTION
func main() {
	var totalUsers int
	var users []User
	var totalRevenue int

	// Counters (real-world analytics)
	var adultCount, teenCount, childCount int

	fmt.Print("Enter number of users: ")
	fmt.Scan(&totalUsers)

	for i := 1; i <= totalUsers; i++ {
		fmt.Printf("\n--- User %d ---\n", i)

		name, age := getUserInput()

		// Validation
		if !validateAge(age) {
			fmt.Println("Invalid age.")
			i--
			continue
		}

		user := User{name: name, age: age}

		// Business logic
		assignTicket(&user)

		// Count categories
		switch user.ticketType {
		case "Adult":
			adultCount++
		case "Teen":
			teenCount++
		case "Child":
			childCount++
		}

		// Revenue
		totalRevenue += user.price

		// Store in slice
		users = append(users, user)
	}

	// Output all users
	printAllUsers(users)

	// Summary (REAL FEATURE 🔥)
	fmt.Println("\n===== Summary =====")
	fmt.Println("Total Users:", len(users))
	fmt.Println("Adults:", adultCount)
	fmt.Println("Teens:", teenCount)
	fmt.Println("Children:", childCount)
	fmt.Println("Total Revenue: PKR", totalRevenue)
	fmt.Println("===================")
}
