package handlers

import (
	"encoding/json"
	"foodapp/kafka"
	"foodapp/models"
	"foodapp/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderHandler struct {
	DB *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{DB: db}
}

func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
	var input struct {
		CustomerName string `json:"customer_name"`
		Address      string `json:"address"`
		Item         string `json:"item"`
		Size         string `json:"size"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	order := models.Order{
		OrderID:      uuid.New().String(),
		CustomerName: input.CustomerName,
		Address:      input.Address,
		Item:         input.Item,
		Size:         input.Size,
		Status:       "PLACED",
	}

	// Save order
	if err := services.CreateOrder(h.DB, &order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	orderBytes, _ := json.Marshal(order)

	// Publish event to Kafka
	if err := kafka.Publish(order.OrderID, string(orderBytes)); err != nil {
		log.Println("Failed to publish Kafka message:", err)
	}

	// Return response
	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *OrderHandler) GetOrderByID(c *fiber.Ctx) error {
	orderID := c.Params("id")

	// fetch order
	var order models.Order
	if err := h.DB.First(&order, "order_id = ?", orderID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	// fetch timeline
	var events []models.OrderEvent
	if err := h.DB.Where("order_id = ?", orderID).Order("timestamp asc").Find(&events).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch order events",
		})
	}

	// response shape
	return c.JSON(fiber.Map{
		"order":    order,
		"timeline": events,
	})
}

func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	var orders []models.Order

	if err := h.DB.Find(&orders).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not fetch orders",
		})
	}

	return c.JSON(fiber.Map{
		"orders": orders,
	})
}
