package main

type Analyzer struct {
	TotalRequests int
	IPCount       map[string]int
	StatusCount   map[int]int
	ClientErrors  int
	ServerErrors  int
}

func NewAnalyzer() *Analyzer {
	analyzer := Analyzer{
		IPCount:     make(map[string]int),
		StatusCount: make(map[int]int),
	}

	return &analyzer
}

func (a *Analyzer) ProcessEntry(entry LogEntry) {
	a.TotalRequests++
	a.IPCount[entry.IPAddress]++
	a.StatusCount[entry.StatusCode]++
	if entry.StatusCode >= 400 && entry.StatusCode <= 499 {
		a.ClientErrors++
	} else if entry.StatusCode >= 500 && entry.StatusCode <= 599 {
		a.ServerErrors++
	}
}
