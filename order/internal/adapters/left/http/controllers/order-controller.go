package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/HlufD/order-ms/internal/adapters/left/http/dto"
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

// @Summary Create a new order
// @Description Create an order based on the provided details
// @Tags orders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderDTO true "Create Order"
// @Success 201 {object} domain.Order
// @Failure 400 {object} shared.ErrorResponse
// @Failure 401 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /orders [post]
func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	token, err := shared.ExtractToken(r)
	if err != nil {
		shared.RespondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	var order domain.Order
	var orderDto dto.CreateOrderDTO

	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = shared.Validate(orderDto)
	if err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	order = orderDto.ToEntity()

	createdOrder, err := oc.orderUseCase.Create(&order, token)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusCreated, createdOrder)
}

// @Summary Get order by ID
// @Description Get details of an order by its ID
// @Tags Orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} domain.Order
// @Failure 404 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /orders/{id} [get]
func (oc *OrderController) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	order, err := oc.orderUseCase.GetOrder(id)
	if err != nil {
		shared.RespondWithError(w, http.StatusNotFound, "Order not found")
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, order)
}

// @Summary Get orders for the authenticated user
// @Description Get a list of orders for the currently authenticated user
// @Tags orders
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Order
// @Failure 401 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /orders/user [get]
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

// @Summary Update an existing order
// @Description Update an order's details by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body dto.UpdateOrderDTO true "Updated Order Data"
// @Success 200 {object} domain.Order
// @Failure 400 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /orders/{id} [put]
func (oc *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var updateOrder domain.UpdateOrder
	var updatedOrderDto dto.UpdateOrderDTO

	if err := json.NewDecoder(r.Body).Decode(&updatedOrderDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := shared.Validate(updatedOrderDto)

	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	updateOrder = updatedOrderDto.ToEntity()

	updatedOrder, err := oc.orderUseCase.UpdateOrder(id, &updateOrder)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, updatedOrder)
}
