package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type LogEntry struct {
	IPAddress    string
	Timestamp    string
	Method       string
	URL          string
	HTTPVersion  string
	StatusCode   int
	ResponseSize int
}

func ParseLog(line string) (LogEntry, error) {
	regex := regexp.MustCompile(`^(\S+) \S+ \S+ \[([^\]]+)\] "(\S+) (\S+) (\S+)" (\d{3}) (\d+)$`)

	result := regex.FindStringSubmatch(line)

	if result == nil {
		return LogEntry{}, fmt.Errorf("invalid log format: %s", line)
	}

	statusCode, err := strconv.Atoi(result[6])
	if err != nil {
		return LogEntry{}, fmt.Errorf("invalid status code: %w", err)
	}

	responseSize, err := strconv.Atoi(result[7])
	if err != nil {
		return LogEntry{}, fmt.Errorf("invalid response size: %w", err)
	}

	log := LogEntry{
		IPAddress:    result[1],
		Timestamp:    result[2],
		Method:       result[3],
		URL:          result[4],
		HTTPVersion:  result[5],
		StatusCode:   statusCode,
		ResponseSize: responseSize,
	}

	return log, nil

}
