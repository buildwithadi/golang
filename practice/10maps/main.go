package main

import (
	"fmt"
	"maps"
)

// map -> hash, object, dict
func main() {
	// creating map
	m := make(map[string]string)

	// adding element
	m["name"] = "golang"
	m["area"] = "backend"

	// getting an element
	fmt.Println(m["name"], m["area"])
	// getting an element which doesn't exists
	fmt.Println(m["i Dont Exist"]) // zeroed value

	n := make(map[string]int)
	n["age"] = 30
	fmt.Println(n["age"])

	// getting length
	fmt.Println(len(m))

	// deleting an element
	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)

	// creating map with elements
	o := map[string]int{"price": 40, "phones": 30}
	fmt.Println(o)

	// ok : boolean : true if value is present
	value, ok := o["phones"]
	fmt.Println(value)
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	//
	m1 := map[string]int{"price": 40, "phone": 20}
	m2 := map[string]int{"price": 50, "phone": 20}

	fmt.Println(maps.Equal(m1, m2))
}
