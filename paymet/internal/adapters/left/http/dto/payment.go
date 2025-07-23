package dto

import "github.com/HlufD/payment-ms/internal/core/domain"

type PaymentDTO struct {
	OrderID       string  `json:"orderId" validate:"required,len=24,hexadecimal"`
	Amount        float64 `json:"amount" validate:"required,gt=0"`
	Status        string  `json:"status" validate:"required,oneof=pending paid failed"`
	Method        string  `json:"method" validate:"required,oneof=credit_card paypal bank_transfer"`
	TransactionID string  `json:"transactionId" validate:"required"`
}

func (dto *PaymentDTO) MapToEntity() *domain.Payment {
	return &domain.Payment{
		OrderID:     dto.OrderID,
		Amount:      dto.Amount,
		Status:      dto.Status,
		Method:      dto.Method,
		Transaction: dto.TransactionID,
	}
}
