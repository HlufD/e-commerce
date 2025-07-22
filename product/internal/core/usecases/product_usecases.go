package usecases

import (
	"fmt"
	"log"

	"github.com/HlufD/products-ms/internal/core/domain"
	"github.com/HlufD/products-ms/internal/core/ports/right/persistence"
)

type ProductUseCase struct {
	productRepository persistence.ProductRepository
}

func NewProductService(productRepository persistence.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository,
	}
}

func (pu *ProductUseCase) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return pu.productRepository.Save(product)
}

func (pu *ProductUseCase) UpdateProduct(id string, product *domain.UpdateProduct) (*domain.Product, error) {
	// check if the product exists
	log.Println(product)

	existingProduct, err := pu.productRepository.GetProductById(id)

	if err != nil {
		return nil, fmt.Errorf("error happened when getting product by id: %w", err)
	}

	if existingProduct == nil {
		return nil, domain.ErrProductNotFound
	}

	return pu.productRepository.Update(id, product)
}

func (pu *ProductUseCase) GetProductById(id string) (*domain.Product, error) {
	product, err := pu.productRepository.GetProductById(id)

	if err != nil {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (pu *ProductUseCase) GetAllProducts() ([]*domain.Product, error) {
	return pu.productRepository.GetAllProducts()
}

func (pu *ProductUseCase) CheckAvailability(id string, quantity int) (bool, error) {
	return pu.productRepository.CheckAvailability(id, quantity)
}
