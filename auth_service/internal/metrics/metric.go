package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"endpoint", "method", "status"},
	)

	HTTPRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "auth_http_request_duration_seconds",
			Help:    "Duration of HTTP requests",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"endpoint", "method"},
	)

	HTTPRequestsInFlight = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "auth_http_requests_in_flight",
			Help: "Number of in-flight HTTP requests",
		},
		[]string{"endpoint"},
	)

	TokensIssuedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "auth_tokens_issued_total",
			Help: "Total number of tokens issued",
		},
	)

	UsersCreatedTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "auth_users_created_total",
			Help: "Total number of users created",
		},
	)
)

func RegisterMetrics() {
	prometheus.MustRegister(
		HTTPRequestsTotal,
		HTTPRequestDuration,
		HTTPRequestsInFlight,
		TokensIssuedTotal,
		UsersCreatedTotal,
	)
}