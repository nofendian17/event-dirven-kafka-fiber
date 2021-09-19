package config

import (
	"github.com/segmentio/kafka-go"
)

func NewKafkaReader(configuration Config) *kafka.Reader {
	kafkaConfig := kafka.ReaderConfig{
		Brokers: []string{
			configuration.Get("KAFKA_HOST"),
		},
		GroupID:  configuration.Get("KAFKA_CONSUMER_GROUP_ID"),
		Topic:    configuration.Get("KAFKA_TOPIC"),
		MinBytes: 10e3,
		MaxBytes: 10e6,
	}
	kafkaReader := kafka.NewReader(kafkaConfig)

	return kafkaReader
}
