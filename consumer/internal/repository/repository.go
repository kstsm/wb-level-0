package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kstsm/wb-level-0/consumer/internal/domain/models"
	"github.com/kstsm/wb-level-0/consumer/internal/mq/dto"
)

type RepositoryI interface {
	SaveOrder(ctx context.Context, order dto.Order) error
	GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error)
	GetAllOrders(ctx context.Context) ([]*models.Order, error)
}

type Repository struct {
	conn *pgxpool.Pool
}

func NewRepository(conn *pgxpool.Pool) RepositoryI {
	return &Repository{
		conn: conn,
	}
}
