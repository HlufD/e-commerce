package dto

import "github.com/HlufD/order-ms/internal/core/domain"

type UpdateOrderDTO struct {
	Status        *string         `json:"status,omitempty" validate:"omitempty,oneof=pending processing shipped completed cancelled"`
	Items         *[]OrderItemDTO `json:"items,omitempty" validate:"omitempty,dive"`
	TotalAmount   *float64        `json:"totalAmount,omitempty" validate:"omitempty,gt=0"`
	PaymentMethod *string         `json:"paymentMethod,omitempty"`
	IsPaid        *bool           `json:"isPaid,omitempty"`
}

func (dto *UpdateOrderDTO) ToEntity() domain.UpdateOrder {
	var items []domain.OrderItem
	if dto.Items != nil {
		items = make([]domain.OrderItem, len(*dto.Items))
		for i, itemDto := range *dto.Items {
			items[i] = itemDto.ToEntity()
		}
	}

	update := domain.UpdateOrder{}

	if dto.Status != nil {
		update.Status = *dto.Status
	}
	if dto.TotalAmount != nil {
		update.TotalAmount = *dto.TotalAmount
	}
	if dto.PaymentMethod != nil {
		update.PaymentMethod = *dto.PaymentMethod
	}
	if dto.IsPaid != nil {
		update.IsPaid = *dto.IsPaid
	}
	if len(items) > 0 {
		update.Items = items
	}

	return update
}
