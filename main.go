package main

import (
	"fmt"
	"foodapp/handlers"
	"foodapp/kafka"
	"foodapp/models"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting Go Food Ordering Service...")

	app := fiber.New()

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not configured")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected successfully")

	err = db.AutoMigrate(
		&models.Order{},
		&models.OrderEvent{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database tables:", err)
	}

	log.Println("Database tables migrated successfully")

	// Initialize Kafka producer
	kafka.InitProducer()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go Food Ordering Service is running")
	})

	orderHandler := handlers.NewOrderHandler(db)

	app.Get("/api/v1/orders/:id", orderHandler.GetOrderByID)
	app.Post("/api/v1/orders", orderHandler.CreateOrder)
	app.Get("/api/v1/orders", orderHandler.GetOrders)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	log.Printf("Server starting on port %s", port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
