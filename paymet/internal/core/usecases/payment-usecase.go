package usecases

import (
	"fmt"

	"github.com/HlufD/payment-ms/internal/core/domain"
	ports "github.com/HlufD/payment-ms/internal/core/ports/right/persistence"
)

type PaymentUseCase struct {
	paymentRepository ports.PaymentRepository
}

func NewPaymentUseCase(paymentRepository ports.PaymentRepository) *PaymentUseCase {
	return &PaymentUseCase{
		paymentRepository: paymentRepository,
	}
}

func (pu *PaymentUseCase) CreatePayment(payment *domain.Payment) (*domain.Payment, error) {

	if payment.Amount <= 0 {
		return nil, fmt.Errorf("payment amount must be greater than zero")
	}

	createdPayment, err := pu.paymentRepository.Create(payment)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment: %w", err)
	}

	return createdPayment, nil
}

func (pu *PaymentUseCase) GetPaymentByID(id string) (*domain.Payment, error) {
	payment, err := pu.paymentRepository.FindByID(id)

	if err != nil {
		return nil, domain.ErrPaymentNotFound
	}

	return payment, nil
}

func (pu *PaymentUseCase) GetPaymentByOrderID(orderID string) (*domain.Payment, error) {
	payment, err := pu.paymentRepository.FindByOrderID(orderID)
	if err != nil {
		return nil, domain.ErrPaymentNotFound
	}

	return payment, nil
}
