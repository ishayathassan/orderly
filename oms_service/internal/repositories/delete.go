package repositories

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/models"
	"time"
)

func Delete(id uint) error {

	start := time.Now()

	err := database.DB.Delete(&models.Order{}, id).Error

	duration := time.Since(start).Seconds()

	metrics.DatabaseQueryDuration.
		WithLabelValues("DELETE").
		Observe(duration)

	return err
}