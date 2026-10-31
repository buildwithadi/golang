package main

// goroutine : lightweight thread of execution managed by Go runtime

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done() // decrements WaitGroup
	fmt.Println("doing task", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1) // increments WaitGroup
		go task(i, &wg)
	}

	wg.Wait()
}
