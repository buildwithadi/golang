package main

import "fmt"

const age = 20

func main() {
	const name string = "golang"

	const (
		multi     = "hello"
		constants = "world"
	)

	fmt.Println(name, age)
	fmt.Println(multi, constants)
}
