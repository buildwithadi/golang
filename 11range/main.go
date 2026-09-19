package main

import "fmt"

// range : iterating over data structures
func main() {
	nums := []int{6, 7, 8}

	// iterating using for loop
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[i], " ")
	}

	// iterating using range : index, value
	sum := 0
	for i, num := range nums {
		fmt.Println(num, i)
		sum = sum + num
	}

	fmt.Println(sum)

	// iteating over map using range
	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
	}

	// iterating string using range
	// starting byte of rune
	// index, unicode
	for i, v := range "golang" {
		fmt.Println(i, v, string(v))
	}

}
