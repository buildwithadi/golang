# Go Log Parser CLI

A lightweight CLI tool built with **Go** to process large Nginx access logs, extract useful information, and generate JSON reports.

Built using Go's standard library, with no external dependencies.

## Tech Stack

- **Language:** Go
- **Libraries:** `os`, `bufio`, `flag`, `fmt`, `time`
- **Input:** Nginx Access Logs
- **Output:** JSON (planned)

## Development Progress

- [x] **Phase 1:** File Reading & Line Counting
- [ ] **Phase 2:** Log Parsing using Regex
- [ ] **Phase 3:** Log Analytics (IPs, Status Codes, Requests)
- [ ] **Phase 4:** JSON Report Generation

## Phase 1 — Performance Results

Successfully processed a **1 GB log file containing over 9 million lines** using sequential, line-by-line reading.

**Command:**

```bash
go run main.go -filename=access.log
```

**Output:**

```text
Number of lines: 9126527
Time take: 1.392491s
```

| Metric | Result |
|---|---|
| File Size | 1 GB |
| Lines Processed | 9,126,527 |
| Execution Time | 1.392491s |

## Getting Started

**1. Clone the repository**

```bash
git clone <repository-url>
cd logCLI
```

**2. Run the application**

```bash
go run main.go -filename=access.log
```

Replace `access.log` with your own log file.

## Next Steps

**Phase 2:** Parse individual Nginx log entries to extract IP addresses, HTTP methods, requested URLs, and status codes using regular expressions.