package models

import "time"


type Order struct {
    ID        uint      `json:"id" example:"1"`

    UserID    string    `json:"user_id" binding:"required" gorm:"not null;index" example:"123"`
    ItemName  string    `json:"item_name" binding:"required" gorm:"type:varchar(100);not null" example:"Laptop"`

    Quantity  int       `json:"quantity" binding:"required" gorm:"not null" example:"2"`
    Amount    float64   `json:"amount" binding:"required" gorm:"type:decimal(10,2);not null" example:"2500.50"`

    Status    string    `json:"status" gorm:"default:'pending'" example:"pending"`
    
    CreatedAt time.Time `json:"created_at" example:"2026-03-08T12:00:00Z"`
    UpdatedAt time.Time `json:"updated_at" example:"2026-03-08T12:00:00Z"`
}

