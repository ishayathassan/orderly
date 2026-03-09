package services

import (
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
)



func GetAllOrders() ([]models.Order, error) {
    return repositories.GetAll()
}