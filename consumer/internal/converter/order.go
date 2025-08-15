package converter

import (
	apimodels "github.com/kstsm/wb-level-0/consumer/internal/api/models"
	"github.com/kstsm/wb-level-0/consumer/internal/domain/models"
)

func ConvertOrderToResponse(in *models.Order) apimodels.GetOrderByIDHandlerResponse {
	delivery := apimodels.Delivery{
		Name:    in.Delivery.Name,
		Phone:   in.Delivery.Phone,
		Zip:     in.Delivery.Zip,
		City:    in.Delivery.City,
		Address: in.Delivery.Address,
		Region:  in.Delivery.Region,
		Email:   in.Delivery.Email,
	}

	payment := apimodels.Payment{
		Transaction:  in.Payment.Transaction,
		RequestID:    in.Payment.RequestID,
		Currency:     in.Payment.Currency,
		Provider:     in.Payment.Provider,
		Amount:       in.Payment.Amount,
		PaymentDT:    in.Payment.PaymentDT,
		Bank:         in.Payment.Bank,
		DeliveryCost: in.Payment.DeliveryCost,
		GoodsTotal:   in.Payment.GoodsTotal,
		CustomFee:    in.Payment.CustomFee,
	}

	items := make([]apimodels.Item, len(in.Items))
	for i, it := range in.Items {
		items[i] = apimodels.Item{
			ChrtID:      it.ChrtID,
			TrackNumber: it.TrackNumber,
			Price:       it.Price,
			Rid:         it.Rid,
			Name:        it.Name,
			Sale:        it.Sale,
			Size:        it.Size,
			TotalPrice:  it.TotalPrice,
			NmID:        it.NmID,
			Brand:       it.Brand,
			Status:      it.Status,
		}
	}

	return apimodels.GetOrderByIDHandlerResponse{
		OrderUID:          in.OrderUID,
		TrackNumber:       in.TrackNumber,
		Entry:             in.Entry,
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            in.Locale,
		InternalSignature: in.InternalSignature,
		CustomerID:        in.CustomerID,
		DeliveryService:   in.DeliveryService,
		ShardKey:          in.ShardKey,
		SmID:              in.SmID,
		DateCreated:       in.DateCreated,
		OofShard:          in.OofShard,
	}
}
