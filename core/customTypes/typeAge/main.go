package main

import "fmt"

type age int

func main() {
	var aman age = 20

	aman.display()
	fmt.Println(aman.isEligible())
	fmt.Println(aman.inMonths())
}

func (a age) display() {
	fmt.Println(a)
}

func (a age) isEligible() bool {
	if a > 18 {
		return true
	} else {
		return false
	}
}

func (a age) inMonths() int {
	return int(a * 12)
}
