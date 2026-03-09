package services

import (
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"orderly/oms-service/internal/utils"
	"time"
)

func UpdateOrder(id uint, updated models.Order) (*models.Order, error) {
	prevOrder, err := repositories.GetByID(id)
	if err != nil {
		return nil, utils.ErrOrderNotFound
	}

	updated.ID = id
	updated.CreatedAt = prevOrder.CreatedAt
	updated.UpdatedAt = time.Now()

	if err := repositories.Update(updated); err != nil {
		return nil, err
	}

	return &updated, nil
}