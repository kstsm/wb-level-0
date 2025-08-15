package mq

import (
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gookit/slog"
	"github.com/kstsm/wb-level-0/consumer/config"
	"github.com/kstsm/wb-level-0/consumer/internal/service"
)

type Consumer struct {
	consumer *kafka.Consumer
	service  service.ServiceI
	topic    string
}

func NewConsumer(cfg config.Config, svc service.ServiceI) (*Consumer, error) {
	conf := &kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers,
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	c, err := kafka.NewConsumer(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	if err := c.Subscribe(cfg.Kafka.Topic, nil); err != nil {
		return nil, fmt.Errorf("failed to subscribe to topic %s: %w", cfg.Kafka.Topic, err)
	}

	return &Consumer{
		consumer: c,
		service:  svc,
		topic:    cfg.Kafka.Topic,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			slog.Error("Kafka error", "error", err)
			continue
		}

		switch *msg.TopicPartition.Topic {
		case "orders":
			if err = c.service.SaveOrder(ctx, msg.Value); err != nil {
				slog.Error("Failed to process order", "error", err)
			}
		}
	}
}

func (c *Consumer) Close() error {
	slog.Info("Closing Kafka consumer")
	return c.consumer.Close()
}
