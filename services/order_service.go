package services

import (
	"encoding/json"
	"foodapp/kafka"
	"foodapp/models"
	"log"
	"time"

	"gorm.io/gorm"
)

var statusFlow = []string{"PLACED", "PREPARING", "COOKING", "OUT_FOR_DELIVERY", "DELIVERED"}

func CreateOrder(db *gorm.DB, order *models.Order) error {
	// Insert the initial order
	if err := db.Create(order).Error; err != nil {
		return err
	}

	// Also insert the initial event (PLACED)
	initialEvent := models.OrderEvent{
		OrderID: order.OrderID,
		Event:   order.Status,
	}
	if err := db.Create(&initialEvent).Error; err != nil {
		return err
	}

	go func(order *models.Order) {
		for i := 1; i < len(statusFlow); i++ {
			time.Sleep(10 * time.Second)

			var statusChecker models.Order
			if err := db.First(&statusChecker, "order_id = ?", order.OrderID).Error; err != nil {
				return
			}

			if statusChecker.Status == "CANCELLED" {
				log.Printf("Order %s status progression stopped because it was cancelled", order.OrderID)
				return
			}

			// Update order status in orders table
			if err := db.Model(&models.Order{}).
				Where("order_id = ?", order.OrderID).
				Update("status", statusFlow[i]).Error; err != nil {
				return // stop if DB fails
			}

			// Insert a new order status event
			newEvent := models.OrderEvent{
				OrderID: order.OrderID,
				Event:   statusFlow[i],
			}
			if err := db.Create(&newEvent).Error; err != nil {
				return
			}

			eventBytes, _ := json.Marshal(newEvent)
			_ = kafka.Publish(order.OrderID, string(eventBytes))
		}
	}(order)

	return nil
}
