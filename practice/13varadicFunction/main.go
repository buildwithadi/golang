package main

import "fmt"

// varadic functions : n number of parameters can be taken

// specific type param
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total = total + v
	}
	return total
}

// any type param
func printValues(values ...interface{}) {
	for _, v := range values {
		fmt.Print(v, " ")
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	printValues(1, "hello", true, "\n")

	// passing slice as parameters
	nums := []int{3, 4, 5, 6}
	fmt.Println(sum(nums...))
}
