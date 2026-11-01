package main

import "fmt"

type Student struct {
	name   string
	rollno int32
	marks  int32
}

func main() {
	harry := Student{
		name:   "Harry",
		rollno: 1,
		marks:  50,
	}

	updateMarks(&harry)

	fmt.Print(harry)
}

func updateMarks(s *Student) {
	s.marks = 100
}
