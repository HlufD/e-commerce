package dto

import (
	"github.com/HlufD/order-ms/internal/core/domain"
)

type CreateOrderDTO struct {
	Items         []OrderItemDTO `json:"items" bson:"items" validate:"required,min=1,dive,required"`
	PaymentMethod string         `json:"paymentMethod" bson:"paymentMethod" validate:"required"`
}

func (dto *CreateOrderDTO) ToEntity() domain.Order {

	items := make([]domain.OrderItem, len(dto.Items))

	for i, itemDto := range dto.Items {
		items[i] = itemDto.ToEntity()
	}

	return domain.Order{
		Items:         items,
		PaymentMethod: dto.PaymentMethod,
	}
}
