package main

import "fmt"

func main() {
	var arr []int = []int{1, 2, 3, 4, 5}

	var res []int = reverse(arr)

	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
}

func reverse(arr []int) []int {

	var n int = len(arr)

	for i := 0; i < n/2; i++ {
		temp := arr[i]
		arr[i] = arr[n-i-1]
		arr[n-i-1] = temp
	}

	return arr
}
