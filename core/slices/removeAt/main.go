package main

import "fmt"

func main() {

	var slice []int = []int{1, 2, 3, 4, 5}

	var res []int = removeAt(slice, 2)

	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}

}

func removeAt(arr []int, index int) []int {
	var l int = len(arr)
	var res []int

	for i := 0; i < l; i++ {
		if i != index {
			res = append(res, arr[i])
		}
	}

	return res
}
