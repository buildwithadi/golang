package main

import (
	"errors"
	"fmt"
)

type Student struct {
	id     int
	name   string
	age    int
	course string
	marks  []int
}

func main() {

	// I only wrote the core functionality of the project which is struct and using struct in the function.

}

// Creating Student Record
func createStudent(id int, name string, age int, course string, marks []int, students []Student) bool {

	stud := Student{
		id:     id,
		name:   name,
		age:    age,
		course: course,
		marks:  marks,
	}

	students = append(students, stud)

	return true
}

// Displaying Student Details
func displayStudent(id int, students []Student) bool {
	var l int = len(students)

	for i := 0; i < l; i++ {
		if students[i].id == id {
			var student Student = students[i]
			fmt.Printf("Id : %d\nName : %s\nAge : %d\nCourse : %s\nMarks : ", student.id, student.name, student.age, student.course)
			for j := 0; j < 5; j++ {
				fmt.Print(student.marks[j], " ")
			}
		}
	}

	return true
}

// Calculate Average
func calculateAverage(id int, students []Student) float32 {
	var student Student

	student, err := findStudent(id, students)
	if err != nil {
		fmt.Print("error", err)
		return 0.0
	}

	var marks []int = student.marks

	var averageMarks float32
	var sum int = 0

	for mark := range marks {
		sum = sum + mark
	}

	averageMarks = float32(sum) / float32(len(marks))

	return averageMarks
}

// Finding Student
func findStudent(id int, students []Student) (Student, error) {

	var l int = len(students)

	for i := 0; i < l; i++ {
		if students[i].id == id {
			return students[i], nil
		}
	}

	return Student{}, errors.New("Not Found")
}
