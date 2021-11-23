package handler

import (
	"encoding/json"
	"net/http"

	"github.com/fizzbuzz-endpoint/internal/metric"
	"github.com/gol4ng/logger"
)

// PrintMetrics shows the top used request with the number of hits for this request
// and the top used values for every parameter.
func PrintMetrics(log logger.LoggerInterface, m *metric.Metric) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")

		m.Lock()
		output := metric.MetricsResponse{
			Number1: m.CNumber1.GetTop(),
			Number2: m.CNumber2.GetTop(),
			Limit:   m.CLimit.GetTop(),
			Str1:    m.CStr1.GetTop(),
			Str2:    m.CStr2.GetTop(),
			Query:   m.CQuery.GetTop(),
		}
		m.Unlock()

		payload, err := json.Marshal(output)
		if err != nil {
			log.Error("faild to marshal metrics : %err%", logger.Error("err", err))
		}

		response.WriteHeader(http.StatusOK)
		response.Write(payload)
	}
}
