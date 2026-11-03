package main

import "fmt"

type Employee struct {
	id   int
	name string
	dept string
}

func main() {
	var employees []Employee = []Employee{
		{id: 1, name: "Karan", dept: "Finance"},
		{id: 2, name: "Rahul", dept: "IT"},
		{id: 3, name: "Priya", dept: "HR"},
		{id: 4, name: "Neha", dept: "Finance"},
		{id: 5, name: "Aman", dept: "IT"},
	}

	m := make(map[string][]string)

	for _, employee := range employees {
		m[employee.dept] = append(m[employee.dept], employee.name)
	}

	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}
