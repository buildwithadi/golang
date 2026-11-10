package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	var count int = 0

	start := time.Now()

	// Accepting filename through CLI
	var filename = flag.String("filename", "temp.log", "read the file")
	flag.Parse()

	// Opening the File
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	// Reading the File Line by Line
	// Count Total Number of lines in the file
	reader := bufio.NewReader(file)
	for {
		_, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		count++
	}

	// Calculating the time taken to read the file
	enlapsed := time.Since(start)
	fmt.Println("Number of lines:", count, "Time take:", enlapsed)

}
