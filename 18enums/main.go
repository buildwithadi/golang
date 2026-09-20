package main

import "fmt"

// enumerated types
type OrderStatus string

// const (
// iota increments the value on go
// 	Received OrderStatus = iota	// 0
// 	Confirmed 					// 1
// 	Prepared 					// 2
// 	Delivered					// 3
// )

const (
	Received  OrderStatus = "Received"
	Confirmed             = "Confirmed"
	Prepared              = "Prepare"
	Delivered             = "Delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Confirmed)
}
