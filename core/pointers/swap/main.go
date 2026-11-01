package main

import "fmt"

func main() {
	var a int
	var b int

	var ptr1 *int = &a
	var ptr2 *int = &b

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Print("Error", err)
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Print("Error", err)
		return
	}

	temp := *ptr1
	*ptr1 = *ptr2
	*ptr2 = temp

	fmt.Print(a, b)
}
