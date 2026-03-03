package models

import "time"


type Order struct {
    ID        uint      `json:"id" gorm:"primaryKey"`

    UserID    string    `json:"user_id" binding:"required" gorm:"not null;index"`
    
	ItemName  string    `json:"item_name" binding:"required" gorm:"type:varchar(100);not null"`
    Quantity  int       `json:"quantity" binding:"required" gorm:"not null"`
    Amount    float64   `json:"amount" binding:"required" gorm:"type:decimal(10,2);not null"`
    
	Status    string    `json:"status" gorm:"default:'pending'"`
    
	CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}