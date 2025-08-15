package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/kstsm/wb-level-0/consumer/internal/domain/models"
	"github.com/kstsm/wb-level-0/consumer/internal/mq/dto"
)

func (s *Service) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error) {
	var cachedOrder models.Order
	cacheKey := fmt.Sprintf("order:%s", orderID.String())

	if err := s.redis.GetJSON(cacheKey, &cachedOrder); err == nil {
		return &cachedOrder, nil
	}

	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	if err := s.redis.SetJSON(cacheKey, order, 0); err != nil {
		slog.Errorf("failed to cache order %s: %v", order.OrderUID, err)
	}

	return order, nil
}

func (s *Service) SaveOrder(ctx context.Context, data []byte) error {
	var order dto.Order
	if err := json.Unmarshal(data, &order); err != nil {
		return fmt.Errorf("failed to unmarshal order: %w", err)
	}

	if err := s.validate.Struct(&order); err != nil {
		slog.Warnf("validation failed for order %s: %v", order.OrderUID, err)
		return err
	}

	err := s.repo.SaveOrder(ctx, order)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("order:%s", order.OrderUID)
	if err := s.redis.SetJSON(cacheKey, order, 0); err != nil {
		slog.Errorf("failed to cache order %s: %v", order.OrderUID, err)
	}

	return nil
}
