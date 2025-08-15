package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/kstsm/wb-level-0/consumer/internal/service"
	"net/http"
)

type HandlerI interface {
	NewRouter() http.Handler
	GetOrderByIDHandler(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.ServiceI
}

func NewHandler(service service.ServiceI) HandlerI {
	return &Handler{
		service: service,
	}
}

func (h Handler) NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/order", h.serveOrderPage)
	r.Get("/order/{order_uid}", h.GetOrderByIDHandler)

	return r
}

func (h Handler) serveOrderPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/order.html")
}
