package metric

import (
	"sync"
)

// Metric struct of counters to count the top used params
type Metric struct {
	CNumber1 Counter
	CNumber2 Counter
	CLimit   Counter
	CStr1    Counter
	CStr2    Counter
	CQuery   Counter

	sync.Mutex
}

// NewMetric initialize counters with top 3 value
func NewMetric() *Metric {
	return &Metric{
		CNumber1: NewCounter(3),
		CNumber2: NewCounter(3),
		CLimit:   NewCounter(3),
		CStr1:    NewCounter(3),
		CStr2:    NewCounter(3),
		CQuery:   NewCounter(3),
	}
}

// PairElement used to count th frequent of parameter's value
type PairElement struct {
	Value string `json:"value"`
	Count int    `json:"count"`
}

// MetricsResponse shows top used request and parameters
type MetricsResponse struct {
	Query   []PairElement `json:"query,omitempty"`
	Number1 []PairElement `json:"number1,omitempty"`
	Number2 []PairElement `json:"number2,omitempty"`
	Limit   []PairElement `json:"limit,omitempty"`
	Str1    []PairElement `json:"str1,omitempty"`
	Str2    []PairElement `json:"str2,omitempty"`
}
