package config

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func NewKafkaWriter(configuration Config) *kafka.Writer {
	_, cancel := NewKafkaWriterContext()
	defer cancel()

	writer := kafka.Writer{
		Addr: kafka.TCP(
			configuration.Get("KAFKA_HOST"),
		),
		Topic:        configuration.Get("KAFKA_TOPIC"),
		Balancer:     &kafka.LeastBytes{},
		WriteTimeout: 10 * time.Second,
		Async: true,
		Logger: nil,
		ErrorLogger: nil,
	}

	return &writer
}

func NewKafkaWriterContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}