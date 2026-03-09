package repositories

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/models"
	"time"
)

func Create(order *models.Order) error {
	start := time.Now()

	err := database.DB.Create(order).Error

	duration := time.Since(start).Seconds()

	metrics.DatabaseQueryDuration.
		WithLabelValues("INSERT").
		Observe(duration)

	return err
}