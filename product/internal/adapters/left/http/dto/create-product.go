package dto

import (
	"time"

	"github.com/HlufD/products-ms/internal/core/domain"
)

type CreateProduct struct {
	Name        string  `json:"name" validate:"required,min=2,max=100"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"required,gte=0"`
	Category    string  `json:"category" validate:"required"`
}

func (cp *CreateProduct) MapToDomainEntity() *domain.Product {
	now := time.Now()
	return &domain.Product{
		Name:        cp.Name,
		Description: cp.Description,
		Price:       cp.Price,
		Stock:       cp.Stock,
		Category:    cp.Category,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
