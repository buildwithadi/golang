package main

import "fmt"

// interfaces (-er)
// it implicitly apply interfaces to the classes which have the provided same method signatures
type paymenter interface {
	pay(amount float32)

	// we can add other methods to it like:
	// refund()
	// currentBalance()
}

// open close principle : open for extension but close for modification

// payment struct
type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	// razorpayPaymentGw := razorpay{} // object of razorpay struct
	// razorpayPaymentGw.pay(amount)

	// correct method:
	p.gateway.pay(amount)
}

// razorpay struct
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("Making payment using razorpay", amount)
}

// strip struct
type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

// paypal struct
type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func main() {
	// newPayment := payment{}
	// newPayment.makePayment(100)

	// correct method:
	// stripePaymentGw := stripe{}
	// paypalPaymentGw := paypal{}
	razorpayPaymentGw := razorpay{}
	newPayment := payment{
		// now just change the object and it will find the specific method automatically
		gateway: razorpayPaymentGw,
	}
	newPayment.makePayment(100)
}
