package dto

import "github.com/HlufD/order-ms/internal/core/domain"

type OrderItemDTO struct {
	ProductID string `json:"productId" bson:"productId" validate:"required"`
	Quantity  int    `json:"quantity" bson:"quantity" validate:"required,gt=0"`
}

func (dto *OrderItemDTO) ToEntity() domain.OrderItem {
	return domain.OrderItem{
		ProductID: dto.ProductID,
		Quantity:  dto.Quantity,
	}
}
