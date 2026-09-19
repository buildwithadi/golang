package main

import (
	"fmt"
	"time"
)

type customer struct {
	name  string
	phone string
}

// order struct
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer            // embedding the struct (inheritance)
}

// constructor
func newOrder(id string, amount float32, status string) *order {

	// initial setup goes here

	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

// we don't need to use '*' here, 'cus we don't need the reference of the obj
func (o order) getAmount() float32 {
	return o.amount
}

// if we don't set any field then the value is zeroed value
// int: 0; float:0.0; string:""; bool: false

func main() {
	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

	fmt.Println("Order Struct", myOrder)

	// adding new value
	myOrder.createdAt = time.Now()

	fmt.Println("Order Struct", myOrder)

	// get specific variable
	fmt.Println(myOrder.status)

	// using method
	myOrder.changeStatus("confirmed")
	fmt.Println(myOrder)

	fmt.Println("Amount:", myOrder.getAmount())

	// creating order from constuructor
	firstOrder := newOrder("1", 30.50, "received")

	fmt.Println(firstOrder.getAmount())

	// creating object
	language := struct {
		name   string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	// struct embedding
	aman := customer{
		name:  "aman",
		phone: "1234567890",
	}

	amanOrder := order{
		id:       "1",
		amount:   100,
		status:   "paid",
		customer: aman,
		// 	customer: customer{
		// 		name: "aditya",
		//		phone: "1234567890",
		//	}
	}

	fmt.Println(amanOrder)
	amanOrder.customer.phone = "11111111"
	fmt.Println(amanOrder)
	fmt.Println(aman)

}
