package main

import "fmt"

// golang only have 'for' for looping
func main() {

	// while loop
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println("i am loop")
	// }

	fmt.Println()

	// for loop
	for i := 0; i <= 5; i++ {
		if i == 2 {
			continue
		}
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println()

	// Range
	for i := range 3 {
		fmt.Println(i)
	}
}
