package middlwares

import (
	"net/http"
	"orderly/auth-service/internal/metrics"
	"time"

	"github.com/gin-gonic/gin"
)

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		endpoint := c.FullPath()
		if endpoint == "" {
			endpoint = "unknown"
		}

		method := c.Request.Method

		// Track in-flight requests
		metrics.HTTPRequestsInFlight.WithLabelValues(endpoint).Inc()
		start := time.Now()

		c.Next()

		duration := time.Since(start).Seconds()
		status := c.Writer.Status()

		// Record metrics
		metrics.HTTPRequestsInFlight.WithLabelValues(endpoint).Dec()
		metrics.HTTPRequestDuration.WithLabelValues(endpoint, method).Observe(duration)
		metrics.HTTPRequestsTotal.WithLabelValues(endpoint, method, http.StatusText(status)).Inc()
	}
}