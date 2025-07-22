package persistence

import "github.com/HlufD/products-ms/internal/core/domain"

type ProductRepository interface {
	Save(product *domain.Product) (*domain.Product, error)
	Update(id string, product *domain.UpdateProduct) (*domain.Product, error)
	GetProductById(id string) (*domain.Product, error)
	GetAllProducts() ([]*domain.Product, error)
	CheckAvailability(id string, quantity int) (bool, error)
}
