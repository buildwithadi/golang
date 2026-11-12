package main

import "sort"

type Analyzer struct {
	TotalRequests int
	IPCount       map[string]int
	StatusCount   map[int]int
	ClientErrors  int
	ServerErrors  int
}

type IPStat struct {
	IP       string
	Requests int
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

func (a *Analyzer) GetTopIPs(limit int) []IPStat {
	res := []IPStat{}

	for i, c := range a.IPCount {
		res = append(res, IPStat{i, c})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Requests > res[j].Requests
	})

	if limit <= 0 {
		return res[:0]
	}

	if limit > len(res) {
		return res
	}

	return res[:limit]
}
