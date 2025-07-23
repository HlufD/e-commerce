package ports

import "github.com/HlufD/order-ms/internal/core/domain"

type OrderRepository interface {
	Create(order *domain.Order) (*domain.Order, error)

	FindByID(id string) (*domain.Order, error)

	FindUserOrder(userId string) ([]*domain.Order, error)

	Update(id string, order *domain.Order) (*domain.Order, error)
}
