# Go Log Analyzer CLI

A lightweight command-line tool written in **Go** to parse and analyze Nginx access logs. It processes log files line by line, extracts useful information, and generates statistics such as total requests, HTTP status distribution, top IP addresses, and client/server errors.

## Features

- **Log Parsing:** Parses Nginx access log entries using regular expressions.
- **Request Statistics:** Calculates the total number of successfully parsed requests.
- **IP Analysis:** Counts requests from each IP address and identifies the top 5 IPs.
- **Status Code Distribution:** Groups requests by HTTP status code.
- **Error Analysis:** Tracks 4xx client errors and 5xx server errors.
- **Malformed Entry Detection:** Counts log entries that don't match the expected format.
- **JSON Report:** Generates a structured JSON report of the analysis.
- **Performance Tracking:** Displays the time taken to process the log file.
- **Memory Efficient:** Reads logs line by line instead of loading the entire file into memory.

## Tech Stack

- **Language:** Go
- **Standard Library:** `bufio`, `flag`, `fmt`, `os`, `regexp`, `strconv`, `encoding/json`, `sort`, `time`

No external Go dependencies are required.

## Project Structure

```text
logCLI/
│
├── main.go             # CLI entry point and file processing
├── parser.go           # Nginx log parsing logic
├── analyzer.go         # Log analysis and statistics
├── analyzer_test.go    # Unit tests
├── models.go           # Data structures (if separated)
├── temp.py             # Generates sample Nginx log files
├── temp.log            # Generated sample log file
└── report.json         # Generated analysis report
```

## Getting Started

### 1. Clone the repository

```bash
git clone <your-repository-url>
cd logCLI
```

### 2. Generate your own log file (Optional)

You can generate a synthetic Nginx access log file using the included `temp.py` script.

Make sure Python is installed, then run:

```bash
python temp.py
```

The script generates a sample log file (`temp.log`) that you can use to test the analyzer.

You can modify the script to generate larger or smaller log files depending on your testing requirements.

### 3. Run the analyzer

Make sure you have Go installed, then execute:

```bash
go run . --filename=temp.log
```

You can also provide a different log file:

```bash
go run . --filename=access.log
```

### 4. Generate the JSON report

The analyzer generates a JSON report containing the processed statistics.

The report is saved as:

```text
report.json
```

## Example Analysis Output

```text
Total Lines: 100000
Successfully Parsed: 99995
Malformed Entries: 5
Unique IPs: 500

Status Distribution:
200 : 85000
301 : 5000
404 : 8000
500 : 1995

Top 5 IP Addresses:
192.168.1.10 - 1500 requests
192.168.1.15 - 1200 requests
192.168.1.20 - 950 requests

Client Errors (4xx): 8000
Server Errors (5xx): 1995

Time Taken: 1.25s
```

*Note: The output above is illustrative. Actual results depend on the input log file.*

## Sample JSON Report

```json
{
  "total_requests": 99995,
  "malformed_entries": 5,
  "unique_ips": 500,
  "status_distribution": {
    "200": 85000,
    "301": 5000,
    "404": 8000,
    "500": 1995
  },
  "client_errors": 8000,
  "server_errors": 1995,
  "top_ips": []
}
```

## Running Tests

Run the unit tests using:

```bash
go test ./...
```

For verbose test output:

```bash
go test -v ./...
```

## Learning Outcomes

This project helped me practice and understand:

- File handling and buffered reading in Go
- Regular expressions and string parsing
- Structs and methods
- Maps and slices
- Sorting data
- Error handling
- Unit testing with Go's `testing` package
- JSON serialization
- CLI argument parsing
- Processing large files sequentially

## Future Improvements

- Support additional log formats.
- Add date-wise and endpoint-wise analytics.
- Support CSV report generation.
- Improve parsing for additional Nginx log formats.
- Explore concurrent processing for very large log files.

## Author

Built as a hands-on Go project to strengthen my understanding of Go fundamentals, file processing, and backend development.
