package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
		line, err := reader.ReadString('\n')

		if len(line) > 0 {
			count++

			line = strings.TrimSpace(line)

			logObj, err := ParseLog(line)

			if err != nil {
				fmt.Print("Error:", err)
				return
			}

			ip := logObj.IPAddress
			timestamp := logObj.Timestamp
			method := logObj.Method
			url := logObj.URL
			statusCode := logObj.StatusCode
			responseSize := logObj.ResponseSize

			fmt.Printf("IP: %s\nTimestamp: %s\nMethod: %s\nURL: %s\nStatus Code: %d\nResponse Size: %d\n\n", ip, timestamp, method, url, statusCode, responseSize)

		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

	}

	// Calculating the time taken to read the file
	fmt.Println()
	enlapsed := time.Since(start)
	fmt.Println("Number of lines:", count, "Time take:", enlapsed)

	defer file.Close()

}
