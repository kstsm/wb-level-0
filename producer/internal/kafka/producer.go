package kafka

import (
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

var errUnknownType = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer
}

func NewProducer() (*Producer, error) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
	}

	p, err := kafka.NewProducer(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	return &Producer{producer: p}, nil
}

func (p *Producer) Produce(message, key []byte, topic string) error {
	kafkaMsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value:     message,
		Key:       key,
		Timestamp: time.Now(),
	}

	kafkaChan := make(chan kafka.Event, 1)

	err := p.producer.Produce(kafkaMsg, kafkaChan)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	e := <-kafkaChan
	switch ev := e.(type) {
	case *kafka.Message:
		if ev.TopicPartition.Error != nil {
			return ev.TopicPartition.Error
		}
		return nil
	case kafka.Error:
		return ev
	default:
		return errUnknownType
	}
}

func (p *Producer) Close() {
	p.producer.Flush(15000)
	p.producer.Close()
}
