package main

import "fmt"

const age = -1

func main() {
	if age >= 18 && age < 100 {
		fmt.Println("Person is Adult")
	} else if age < 18 && age > 0 {
		fmt.Println("Person is not an Adult")
	} else {
		fmt.Println("Invalid Person age")
	}

	var role = "admin"
	var hasPermission = false

	if role == "admin" || hasPermission {
		fmt.Println("Something is true")
	}

	if role == "admin" && hasPermission {
		fmt.Println("Something is not true")
	}

	// we can declare variable inside if construct
	if age := 15; age <= 19 {
		fmt.Println("He is a teenager")
	}

	// go does not have ternary operator, use normal if else
}
