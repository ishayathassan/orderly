package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (

	// Total HTTP requests
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "oms_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "endpoint", "status"},
	)

	// Request latency
	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "oms_http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "endpoint"},
	)

	// Active requests
	HTTPRequestsInFlight = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "oms_http_requests_in_flight",
			Help: "Number of in-flight HTTP requests",
		},
		[]string{"endpoint"},
	)

	// Business metric: orders created
	OrdersCreatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "oms_orders_created_total",
			Help: "Total number of orders created",
		},
	)

	// Database query latency
	DatabaseQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "oms_database_query_duration_seconds",
			Help:    "Database query execution time",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"query_type"},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(
		HTTPRequestsTotal,
		HTTPRequestDuration,
		HTTPRequestsInFlight,
		OrdersCreatedTotal,
		DatabaseQueryDuration,
	)
}