package dto

import (
	"github.com/HlufD/products-ms/internal/core/domain"
)

type UpdateProduct struct {
	Name        *string  `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Description *string  `json:"description,omitempty"`
	Price       *float64 `json:"price,omitempty" validate:"omitempty,gt=0"`
	Stock       *int     `json:"stock,omitempty" validate:"omitempty,gte=0"`
	Category    *string  `json:"category,omitempty"`
}

func (up *UpdateProduct) MapToDomainEntity() *domain.UpdateProduct {

	updateProduct := &domain.UpdateProduct{}

	if up.Name != nil {
		updateProduct.Name = *up.Name
	}
	if up.Description != nil {
		updateProduct.Description = *up.Description
	}
	if up.Price != nil {
		updateProduct.Price = *up.Price
	}
	if up.Stock != nil {
		updateProduct.Stock = *up.Stock
	}
	if up.Category != nil {
		updateProduct.Category = *up.Category
	}

	return updateProduct
}
