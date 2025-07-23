package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	adapters "github.com/HlufD/payment-ms/internal/adapters/left/http"
	"github.com/HlufD/payment-ms/internal/adapters/left/http/dto"
	"github.com/HlufD/payment-ms/internal/core/domain"
	"github.com/HlufD/payment-ms/internal/core/usecases"
	"github.com/HlufD/payment-ms/shared"
	"github.com/go-chi/chi/v5"
)

type PaymentController struct {
	paymentUseCase *usecases.PaymentUseCase
}

func NewPaymentController(paymentUseCase *usecases.PaymentUseCase) *PaymentController {
	return &PaymentController{
		paymentUseCase: paymentUseCase,
	}
}

// MakePayment godoc
// @Summary Create a new payment
// @Description Makes a payment and updates order status to success
// @Tags Payments
// @Accept json
// @Produce json
// @Param payment body dto.PaymentDTO true "Payment request body"
// @Success 200 {object} domain.Payment
// @Failure 400 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /payments [post]
func (pc *PaymentController) Routes(r chi.Router) http.Handler {
	return r.Route("/payments", func(r chi.Router) {
		r.Post("/", pc.MakePayment)
		r.Get("/{id}", pc.GetPaymentByID)
	})
}

func (pc *PaymentController) MakePayment(w http.ResponseWriter, r *http.Request) {
	var payment *domain.Payment
	var paymentDto dto.PaymentDTO

	if err := json.NewDecoder(r.Body).Decode(&paymentDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := shared.Validate(paymentDto); err != nil {
		shared.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	payment = paymentDto.MapToEntity()

	result, err := pc.paymentUseCase.CreatePayment(payment)
	if err != nil {
		shared.RespondWithError(w, http.StatusInternalServerError, "Failed to create payment")
		return
	}

	baseURL := os.Getenv("ORDER_SERVICE_URL")
	if baseURL == "" {
		shared.RespondWithError(w, http.StatusInternalServerError, "Order service URL not set")
		return
	}

	patchBody := map[string]any{
		"status": "success",
		"isPaid": true,
	}

	orderPath := fmt.Sprintf("/%s", strings.TrimSpace(payment.OrderID))
	log.Println(baseURL + orderPath)

	httpClient := adapters.NewHttpClient(baseURL, 5*time.Second)
	if err := httpClient.Put(r.Context(), orderPath, patchBody, nil); err != nil {
		log.Println(err)
		shared.RespondWithError(w, http.StatusInternalServerError, "Failed to update order status")
		return
	}

	shared.RespondWithJSON(w, http.StatusOK, result)
}

// GetPaymentByID godoc
// @Summary Get a payment by ID
// @Description Returns payment information by payment ID
// @Tags Payments
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} domain.Payment
// @Failure 404 {object} shared.ErrorResponse
// @Failure 500 {object} shared.ErrorResponse
// @Router /payments/{id} [get]
func (pc *PaymentController) GetPaymentByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	payment, err := pc.paymentUseCase.GetPaymentByID(id)
	log.Println(err)

	if err != nil {
		if err == domain.ErrPaymentNotFound {
			shared.RespondWithError(w, http.StatusNotFound, err.Error())
			return
		} else {
			shared.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch payment")
			return
		}
	}

	shared.RespondWithJSON(w, http.StatusOK, payment)
}
