package main

import "fmt"

// arrays : numbered sequence of specific length
// - fixed size, that is predictable
// - memory optimization
// - constant time access
func main() {
	// zeroed values
	// int -> 0; bool -> false; float -> 0.0; string -> empty

	var nums [5]int

	// length
	fmt.Println("length:", len(nums))

	nums[0] = 1
	fmt.Println(nums[0])

	fmt.Println(nums)

	// declare array in single line
	values := [3]int{1, 2, 3}
	fmt.Println(values)

	// 2d array
	arr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr)

}
