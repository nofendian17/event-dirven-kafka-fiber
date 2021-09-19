package repository

import (
	"github.com/segmentio/kafka-go"
	"kafka-fibre/config"
	"kafka-fibre/entitiy"
)

func NewProduceRepository(kafkaWriter *kafka.Writer) ProduceRepository {
	return &produceRepositoryImpl{
		KafkaWriter: kafkaWriter,
	}
}

type produceRepositoryImpl struct {
	KafkaWriter *kafka.Writer
}

func (repository *produceRepositoryImpl) Produce(event *entitiy.Event) error {
	ctx, cancel := config.NewKafkaWriterContext()
	defer cancel()

	err := repository.KafkaWriter.WriteMessages(
		ctx,
		kafka.Message{
			Key:   []byte(event.Key),
			Value: []byte(event.Value),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
