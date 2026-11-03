package main

import "fmt"

func main() {
	var arr []int = []int{2, 4, 2, 3, 4}

	var res int = firstDuplicate(arr)

	if res == -1 {
		fmt.Print("No duplicates found")
	} else {
		fmt.Print(res)
	}
}

func firstDuplicate(arr []int) int {
	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if m[arr[i]] == 0 {
			m[arr[i]] = 1
		} else {
			m[arr[i]]++
		}
	}

	for key, value := range m {
		if value == 2 {
			return key
		}
	}

	return -1
}
