package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
)

var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of incoming requests",
	},
	[]string{"path"})

var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "HTTP response status codes",
	},
	[]string{"status"})

var latency = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_response_time_seconds",
		Help: "Latency of HTTP requests",
	}, []string{"path"})
