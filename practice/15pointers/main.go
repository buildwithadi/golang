package main

import (
	"fmt"
)

func changeNum(n int) {
	n = 5
	fmt.Println("Changed Number", n)
}

func changeNumPntr(ptr *int) {
	*ptr = 5 // store 5 in the pointer location
	fmt.Println("Changed Number", *ptr)
}

func main() {
	// trying to change number using function
	n := 1
	changeNum(n)
	fmt.Println("After changeNumber function, n in main", n) // failed to change the n

	// changing number using pointers
	fmt.Println("Memory address", &n)
	changeNumPntr(&n)
	fmt.Println("After changeNumPntr function, n in main", n) // success to change the n

}
