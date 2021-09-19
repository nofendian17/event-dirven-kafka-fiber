package worker

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConsumer interface {
	Consume()
}

func NewKafkaConsumer(kafkaReader *kafka.Reader) KafkaConsumer {
	return &consumer{KafkaReader: kafkaReader}
}

type consumer struct {
	KafkaReader *kafka.Reader
}

func (c *consumer) Consume() {
	ctx := context.Background()
	for {
		msg, err := c.KafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
