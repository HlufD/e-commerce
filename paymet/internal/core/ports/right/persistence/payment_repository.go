package ports

import "github.com/HlufD/payment-ms/internal/core/domain"

type PaymentRepository interface {
	Create(payment *domain.Payment) (*domain.Payment, error)

	FindByID(id string) (*domain.Payment, error)

	FindByOrderID(orderID string) (*domain.Payment, error)
}
