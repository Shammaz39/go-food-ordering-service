package kafka

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

var writer *kafka.Writer

func InitProducer() {
	broker := os.Getenv("KAFKA_BROKER")
	topic := os.Getenv("KAFKA_TOPIC")

	if broker == "" {
		broker = "localhost:9092"
	}

	if topic == "" {
		topic = "food-orders"
	}

	writer = &kafka.Writer{
		Addr:     kafka.TCP(broker),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	log.Println("Kafka producer initialized")
}

func Publish(key, value string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(value),
		Time:  time.Now(),
	}

	return writer.WriteMessages(context.Background(), msg)
}
