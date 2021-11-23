package service

import (
	"fmt"
	"net/http"

	"github.com/fizzbuzz-endpoint/internal/metric"
)

// GetMetric intializes once and returns Metric object
func (container *Container) GetMetric() *metric.Metric {
	if container.metric == nil {
		container.metric = metric.NewMetric()
	}
	return container.metric
}

// MonitoringMiddleware to calculate metrics
func (container *Container) MonitoringMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		values := r.URL.Query()
		if len(values) != 0 {
			m := container.GetMetric()
			m.Lock()
			m.CNumber1.Inc(values.Get("int1"))
			m.CNumber2.Inc(values.Get("int2"))
			m.CLimit.Inc(values.Get("limit"))
			m.CStr1.Inc(values.Get("str1"))
			m.CStr2.Inc(values.Get("str2"))
			m.CQuery.Inc(values.Encode())
			m.Unlock()

			container.logger.Info(fmt.Sprintf("monitoring: increment counters by one to values: int1:%s - int2:%s - limit:%s - str1:%s - str2:%s",
				values.Get("int1"), values.Get("int2"), values.Get("limit"), values.Get("str1"), values.Get("str2")))
		}

		h.ServeHTTP(w, r)
	})
}
