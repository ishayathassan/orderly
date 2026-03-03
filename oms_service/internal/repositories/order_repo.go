package repositories

import (
	"orderly/oms-service/internal/database"
	"orderly/oms-service/internal/models"
)

func GetAll() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Find(&orders).Error
	return orders, err
}

func GetByID(id uint) (models.Order, error) {
	var order models.Order
	err := database.DB.First(&order, id).Error
	return order, err
}

func Create(order *models.Order) error {
	return database.DB.Create(order).Error
}

func Delete(id uint) error {
	return database.DB.Delete(&models.Order{}, id).Error
}

func Update(order models.Order) error {
	return database.DB.Save(order).Error
}