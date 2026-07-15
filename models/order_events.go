package models

import (
	"time"

	"gorm.io/datatypes"
)

type OrderEvent struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   string         `gorm:"not null;index" json:"order_id"`
	Event     string         `gorm:"type:varchar(50);not null" json:"event"`
	Timestamp time.Time      `gorm:"autoCreateTime" json:"timestamp"`
	Meta      datatypes.JSON `gorm:"type:jsonb" json:"meta"`
}
