package models

import "time"

type Order struct {
	ID uint `json:"id"`

	UserID string `json:"user_id" binding:"required"`

	ItemName string  `json:"item_name" binding:"required"`
	Quantity int     `json:"quantity" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`

	Status string `json:"status" default:"pending"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
