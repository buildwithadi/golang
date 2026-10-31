package main

import (
	"fmt"
	"math/rand"
	"time"
)

// channels : used to communicate between two go routines
// is single go routine then channels are prefered, else WaitGroup

// sending data
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("process", num)
		time.Sleep(time.Second)
	}
}

// receiving data
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// unbuffer channel
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

// buffered channel
// <-chan : send only chan; chan<- : receive only chan
func emailSender(emailChan <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailChan {
		fmt.Println("sending email to", email)
	}
}

func main() {

	// sending data
	numChan := make(chan int)
	go processNum(numChan)
	i := 0
	for {
		numChan <- rand.Intn(100)
		i += 1
		if i == 5 {
			break
		}
	}

	// receiving data
	result := make(chan int)
	go sum(result, 4, 5)
	res := <-result // block
	fmt.Println(res)

	// unbuffer channel
	done := make(chan bool)
	go task(done)
	<-done // block

	// buffered channel
	emailChan := make(chan string, 100)
	done2 := make(chan bool)

	go emailSender(emailChan, done2)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
		time.Sleep(time.Second)
	}

	fmt.Println("done sending.")
	close(emailChan) // built-in cline function
	<-done2

	// multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2", chan2Val)
		}
	}

	// emailChan <- "1@example.com"
	// emailChan <- "2@example.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	<-done

	// messageChan := make(chan string)
	// messageChan <- "ping"
	// msg := <-messageChan
	// fmt.Println(msg)
}
