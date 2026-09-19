package main

import (
	"fmt"
)

// name (params) return_type
func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, bool, string, int) {
	return "golang", false, "typescript", 7
}

func processIt(fn func(a int) int) int {
	return fn(1) + 3
}

func processedIt() func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	result := add(2, 3)
	fmt.Println(result)

	fmt.Println(getLanguages())

	fn := func(a int) int {
		return 2
	}

	fmt.Println(processIt(fn))

	fmt.Println(processedIt())
}
