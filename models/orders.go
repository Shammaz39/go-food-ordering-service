package models

import "time"

type Order struct {
	OrderID      string    `gorm:"primaryKey;unique;not null" json:"order_id"`
	CustomerName string    `gorm:"not null" json:"customer_name"`
	Address      string    `gorm:"not null" json:"address"`
	Item         string    `gorm:"not null" json:"item"`
	Size         string    `gorm:"not null" json:"size"`
	Status       string    `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
