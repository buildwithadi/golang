package main

import (
	"fmt"
	"slices"
)

// slice : dynamic arrays
// most used construct in go
// useful methods
func main() {
	// uninitialize slice in nil
	var nums []int

	fmt.Println(nums)
	fmt.Println(len(nums))

	//   make(datatype, initializing capacity, array capacity)
	var slice = make([]int, 2, 5)
	fmt.Println("capacity:", cap(slice))
	fmt.Println(slice)

	slice = append(slice, 1)
	fmt.Println(slice)

	// copy function
	var values = make([]int, 2, 5)
	var values2 = make([]int, len(nums))

	values = append(values, 2)

	copy(values2, values)

	fmt.Println(values2)

	// slice operator
	var slice2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice2[0:2])
	fmt.Println(slice2[2:])

	// slice package
	var a = []int{1, 2, 3}
	var b = []int{1, 2, 3}

	fmt.Println(slices.Equal(a, b))

	var arr2 = [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(arr2)
}
