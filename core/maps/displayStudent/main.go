package main

import "fmt"

func main() {
	students := map[int][]int{
		1: []int{80, 70, 75},
		2: []int{62, 60, 65},
		3: []int{71, 57, 78},
		4: []int{83, 90, 95},
		5: []int{80, 77, 65},
	}

	display(students, 2)
}

func display(student map[int][]int, id int) bool {
	for key, value := range student {
		if key == id {
			fmt.Print(key, ":", value)
		}
	}

	return true
}
