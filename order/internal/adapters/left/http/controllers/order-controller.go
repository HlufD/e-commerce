package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HlufD/order-ms/internal/core/domain"
	"github.com/HlufD/order-ms/internal/core/usecases"
	"github.com/HlufD/order-ms/shared"

	"github.com/go-chi/chi/v5"
)

type OrderController struct {
	orderUseCase usecases.OrderUseCase
}

func NewOrderController(orderUseCase usecases.OrderUseCase) *OrderController {
	return &OrderController{
		orderUseCase: orderUseCase,
	}
}

func (oc *OrderController) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/", oc.CreateOrder)
	r.Get("/{id}", oc.GetOrderByID)
	r.Get("/user", oc.GetUserOrders)
	r.Put("/{id}", oc.UpdateOrder)

	return r
}

func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	token, err := shared.ExtractToken(r)
	if err != nil {
		shared.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	createdOrder, err := oc.orderUseCase.Create(&order, token)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusCreated, createdOrder)
}

func (oc *OrderController) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	order, err := oc.orderUseCase.GetOrder(id)
	if err != nil {
		shared.RespondWithError(w, http.StatusNotFound, "Order not found")
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, order)
}

func (oc *OrderController) GetUserOrders(w http.ResponseWriter, r *http.Request) {
	token, err := shared.ExtractToken(r)
	if err != nil {
		shared.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	orders, err := oc.orderUseCase.GetUserOrders(token)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, orders)
}

func (oc *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updateOrder domain.UpdateOrder
	if err := json.NewDecoder(r.Body).Decode(&updateOrder); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	updatedOrder, err := oc.orderUseCase.UpdateOrder(id, &updateOrder)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, updatedOrder)
}
