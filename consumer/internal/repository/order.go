package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/jackc/pgx/v5"
	"github.com/kstsm/wb-level-0/consumer/internal/apperrors"
	"github.com/kstsm/wb-level-0/consumer/internal/domain/models"
	"github.com/kstsm/wb-level-0/consumer/internal/mq/dto"
)

func (r *Repository) GetAllOrders(ctx context.Context) ([]*models.Order, error) {
	var data []*models.Order

	err := r.conn.QueryRow(ctx, GetAllOrders).Scan(&data)
	if err != nil {
		return nil, fmt.Errorf("QueryRow-GetAllOrders: %w", err)
	}

	return data, nil
}

func (r *Repository) GetOrderByID(ctx context.Context, orderID uuid.UUID) (*models.Order, error) {
	var data *models.Order

	err := r.conn.QueryRow(ctx, GetOrderByID, orderID).Scan(&data)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrOrderNotFound
		}

		return nil, fmt.Errorf("QueryRow-GetOrderByID: %w", err)
	}

	return data, nil
}

func (r *Repository) SaveOrder(ctx context.Context, order dto.Order) error {
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("BeginTx: %w", err)
	}

	defer func() {
		if err != nil {
			err := tx.Rollback(ctx)
			if err != nil {
				slog.Errorf("repo.SaveOrder: tx.Rollback")
			}
		} else {
			err := tx.Commit(ctx)
			if err != nil {
				slog.Errorf("repo.SaveOrder: tx.Commit")
			}
		}
	}()

	var orderID int
	err = tx.
		QueryRow(ctx, SaveOrder,
			order.OrderUID,
			order.TrackNumber,
			order.Entry,
			order.Locale,
			order.InternalSignature,
			order.CustomerID,
			order.DeliveryService,
			order.ShardKey,
			order.SmID,
			order.DateCreated,
			order.OofShard,
		).
		Scan(&orderID)
	if err != nil {
		return fmt.Errorf("QueryRow-SaveOrder: %w", err)
	}

	_, err = tx.Exec(ctx, SaveDelivery,
		orderID,
		order.Delivery.Name,
		order.Delivery.Phone,
		order.Delivery.Zip,
		order.Delivery.City,
		order.Delivery.Address,
		order.Delivery.Region,
		order.Delivery.Email,
	)
	if err != nil {
		return fmt.Errorf("QueryRow-SaveDelivery: %w", err)
	}

	_, err = tx.Exec(ctx, SavePayment,
		orderID,
		order.Payment.Transaction,
		order.Payment.RequestID,
		order.Payment.Currency,
		order.Payment.Provider,
		order.Payment.Amount,
		order.Payment.PaymentDT,
		order.Payment.Bank,
		order.Payment.DeliveryCost,
		order.Payment.GoodsTotal,
		order.Payment.CustomFee,
	)
	if err != nil {
		return fmt.Errorf("QueryRow-SavePayment: %w", err)
	}

	for _, item := range order.Items {
		_, err = tx.Exec(ctx, SaveItem,
			orderID,
			item.ChrtID,
			item.TrackNumber,
			item.Price,
			item.Rid,
			item.Name,
			item.Sale,
			item.Size,
			item.TotalPrice,
			item.NmID,
			item.Brand,
			item.Status,
		)
		if err != nil {
			return fmt.Errorf("QueryRow-SaveItem: %w", err)
		}
	}

	return nil
}
