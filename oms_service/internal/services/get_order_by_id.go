package services

import (
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"orderly/oms-service/internal/utils"
)

func GetOrderByID(id uint) (*models.Order, error) {
	order, err := repositories.GetByID(id)
	if err != nil {
		return nil, utils.ErrOrderNotFound
	}
	return &order, nil
}