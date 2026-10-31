package main

import (
	"fmt"
)

// generic type
// any == interface{} : both works same
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

func printSliceScoped[T int | string](items []T) {
	for _, item := range items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}

// generics in struct

type stack[T any] struct {
	elements []T
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	names := []string{"Aditya", "Aman", "Abhishek"}
	printSlice(names)
	printSlice(nums)

	booleans := []bool{true, false, true, false}
	printSlice(booleans)
	// printSliceScoped(booleans) -> this will give error; cus printSliceScoped is scoped to int and string only

	// generics in struct
	myString := stack[string]{
		elements: []string{"golang", "python", "typescript"},
	}

	myInteger := stack[int]{
		elements: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(myString)
	fmt.Println(myInteger)
}
