package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Print("Array 1 (before changing): ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()

	var check []int = arr

	arr[2] = 99

	fmt.Print("Array 1 (after changing): ")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()

	fmt.Print("Array 2: ")
	for i := 0; i < len(check); i++ {
		fmt.Print(check[i], " ")
	}
}
