package main

import (
	"fmt"
	"math"
)

func main() {
	var value int
	var max int = math.MinInt
	var ptr *int = &max
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Print("Error", err)
		return
	}

	var arr []int

	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Print("Error", err)
			return
		}

		arr = append(arr, value)

		if value > *ptr {
			*ptr = value
		}
	}

	fmt.Print(max)

}
