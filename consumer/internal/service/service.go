package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/kstsm/wb-level-0/consumer/internal/cache"
	"github.com/kstsm/wb-level-0/consumer/internal/domain/models"
	"github.com/kstsm/wb-level-0/consumer/internal/repository"
)

type ServiceI interface {
	SaveOrder(ctx context.Context, raw []byte) error
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error)
}

type Service struct {
	repo     repository.RepositoryI
	validate *validator.Validate
	redis    *cache.Redis
}

func NewService(repo repository.RepositoryI, validate *validator.Validate, redisClient *cache.Redis) *Service {
	return &Service{
		repo:     repo,
		validate: validate,
		redis:    redisClient,
	}
}
