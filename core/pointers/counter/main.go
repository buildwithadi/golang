package main

import "fmt"

func main() {
	var count int = 0
	var ptr *int = &count
	var choice int = 1

	for choice != 4 {
		fmt.Printf("\n\nCounter: %d\n\n", *ptr)

		fmt.Printf("1. Increment\n2. Decrement\n3. Reset\n4. Exit\nEnter your choice: ")

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Print("Error", err)
			return
		}

		switch choice {
		case 1:
			increment(ptr)
		case 2:
			decrement(ptr)
		case 3:
			reset(ptr)
		case 4:
			break
		default:
			fmt.Print("Enter a valid Choice")
		}
	}
}

func increment(ptr *int) {
	*ptr = *ptr + 1
}

func decrement(ptr *int) {
	if *ptr-1 >= 0 {
		*ptr = *ptr - 1
	}
}

func reset(ptr *int) {
	*ptr = 0
}
