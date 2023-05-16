package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "auth_request_total",
}, []string{"status", "method"})

var DurationHistogram = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "auth_request_duration",
}, []string{"method"})
