package repositories

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/metrics"
	"orderly/oms-service/internal/models"
	"time"
)


func Update(order models.Order) error {

    start := time.Now()

    err := database.DB.Save(&order).Error

    duration := time.Since(start).Seconds()

    metrics.DatabaseQueryDuration.
        WithLabelValues("UPDATE").
        Observe(duration)

    return err
}