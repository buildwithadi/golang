package main

import (
	"fmt"
	"math"
)

func main() {

	var arr []int
	var choice int = 1

	for choice != 2 {
		fmt.Printf("1. Add\n2. Exit\nEnter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Print("Error", err)
			return
		}

		fmt.Println()

		switch choice {
		case 1:
			var value int
			fmt.Print("Enter a number: ")
			_, err := fmt.Scan(&value)
			if err != nil {
				fmt.Print("Error", err)
				return
			}
			arr = append(arr, value)
		}

		fmt.Println()
	}

	fmt.Print("Display: ")
	display(arr)
	fmt.Println("Sum:", sum(arr))
	fmt.Println("Average:", average(arr))
	fmt.Println("Minimum:", minimum(arr))
	fmt.Println("Maximum:", maximum(arr))
}

func display(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()
}

func sum(arr []int) int {

	var sum int = 0

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}

	return sum
}

func average(arr []int) float64 {
	return float64(sum(arr)) / float64(len(arr))
}

func minimum(arr []int) int {
	var min int = math.MaxInt

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}

func maximum(arr []int) int {
	var max int = math.MinInt

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
