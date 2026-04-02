package main

import "fmt"

func main() {

	// 1.Basic variable full syntax
	var name string = "Aafia"

	// 2.Short syntax most used syntax
	age := 23

	// 3.Multiple variable declaration
	a, b := 8, 23

	// 4.Default value
	var num int
	var text string
	var check bool

	// 5.Constant (cannot be changed)
	const country = "Pakistan"
	fmt.Println(country)

	// 6.Type inference
	city := "Lahore"
	price := 999.99
	fmt.Println(city, price)

	// Output
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Sum of a and b:", a+b)

	fmt.Println("Default int:", num)
	fmt.Println("Default string:", text)
	fmt.Println("Default bool:", check)

	fmt.Println("Constant:", country)
	fmt.Println("Type inference:", city, price)

}
