package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gookit/slog"
	"github.com/kstsm/wb-level-0/consumer/internal/apperrors"
	"github.com/kstsm/wb-level-0/consumer/internal/converter"
	"github.com/kstsm/wb-level-0/consumer/internal/utils"
	"net/http"
)

func (h Handler) GetOrderByIDHandler(w http.ResponseWriter, r *http.Request) {
	orderIDParam := chi.URLParam(r, "order_uid")

	orderID, err := uuid.Parse(orderIDParam)
	if err != nil {
		slog.Warnf("Invalid order UID format: %s", orderIDParam)
		utils.WriteError(w, http.StatusBadRequest, "неверный формат order_uid")
		return
	}

	result, err := h.service.GetOrderByID(r.Context(), orderID)
	if err != nil {
		switch {
		case errors.Is(err, apperrors.ErrOrderNotFound):
			slog.Warnf("Order not found: %s", orderID)
			utils.WriteError(w, http.StatusNotFound, "заказ не найден")
		default:
			slog.Errorf("Failed to get order %s: %v", orderID, err)
			utils.WriteError(w, http.StatusInternalServerError, "внутренняя ошибка сервера")
		}

		return
	}

	response := converter.ConvertOrderToResponse(result)
	utils.SendJSON(w, http.StatusOK, response)
}
