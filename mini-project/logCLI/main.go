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
	var analyzer *Analyzer = NewAnalyzer()
	var malformedEntries int = 0

	start := time.Now()

	// Accepting filename through CLI
	var filename = flag.String("filename", "temp.log", "read the file")
	flag.Parse()

	// Opening the File
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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
				malformedEntries++

			} else {
				// storing data in analyzer struct
				analyzer.ProcessEntry(logObj)
			}

		}

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

	}

	// Analyzing Result:
	topIPs := analyzer.GetTopIPs(10)

	fmt.Println("===================")
	fmt.Println("===GO LOG PARSER===")
	fmt.Println("===================")

	fmt.Println("\nTotal Lines:", count)
	fmt.Println("Sucessfully Parsed:", analyzer.TotalRequests)
	fmt.Println("Malformed Entries:", malformedEntries)

	fmt.Println("\nHTTP Status Distribution:")
	for status, count := range analyzer.StatusCount {
		fmt.Println(status, ":", count)
	}

	fmt.Println("\nTop 10 IP Addresses:")
	for _, data := range topIPs {
		fmt.Println(data.IP, "-", data.Requests, "requests")
	}

	fmt.Println("\nClient Errors (4xx):", analyzer.ClientErrors)
	fmt.Println("Server Errors (5xx):", analyzer.ServerErrors)

	// Calculating the time taken to read the file
	fmt.Println()
	enlapsed := time.Since(start)
	fmt.Println("Processing Time:", enlapsed)

}
