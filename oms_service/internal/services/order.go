package services

import (
	"errors"
	"orderly/oms-service/internal/models"
	"orderly/oms-service/internal/repositories"
	"time"
)

var ErrOrderNotFound = errors.New("order not found")

func CreateOrder(order models.Order) (*models.Order, error) {
	now := time.Now()
    order.CreatedAt = now
    order.UpdatedAt = now

    if err := repositories.Create(&order); err != nil {
        return nil, err
    }
    return &order, nil
}

func GetOrderByID(id uint) (*models.Order, error) {
    order, err := repositories.GetByID(id)
    if err != nil {
        return nil, ErrOrderNotFound
    }
    return &order, nil
}

func DeleteOrder(id uint) error {
    _, err := repositories.GetByID(id)
    if err != nil {
        return ErrOrderNotFound
    }
    return repositories.Delete(id)
}

func UpdateOrder(id uint, updated models.Order) (*models.Order, error) {
    prevOrder, err := repositories.GetByID(id)
    if err != nil {
        return nil, ErrOrderNotFound
    }

    updated.ID = id
    updated.CreatedAt = prevOrder.CreatedAt
    updated.UpdatedAt = time.Now()

    if err := repositories.Update(updated); err != nil {
        return nil, err
    }

    return &updated, nil
}

func GetAllOrders() ([]models.Order, error) {
    return repositories.GetAll()
}