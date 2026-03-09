package middlewares

import (
	"strconv"
	"time"

	"orderly/oms-service/internal/metrics"

	"github.com/gin-gonic/gin"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method
		endpoint := c.FullPath()

		// Track active requests
		metrics.HTTPRequestsInFlight.WithLabelValues(endpoint).Inc()

		start := time.Now()

		// Process request
		c.Next()

		duration := time.Since(start).Seconds()
		status := strconv.Itoa(c.Writer.Status())

		// Record metrics
		metrics.HTTPRequestsTotal.
			WithLabelValues(method, endpoint, status).
			Inc()

		metrics.HTTPRequestDuration.
			WithLabelValues(method, endpoint).
			Observe(duration)

		metrics.HTTPRequestsInFlight.
			WithLabelValues(endpoint).
			Dec()
	}
}