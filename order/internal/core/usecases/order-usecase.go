package usecases

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/HlufD/order-ms/internal/core/domain"
	httpports "github.com/HlufD/order-ms/internal/core/ports/left/http"
	databaseports "github.com/HlufD/order-ms/internal/core/ports/right/persistence"
	"github.com/HlufD/order-ms/shared"
)

type OrderUseCase struct {
	httpClient      httpports.HttpRequester
	orderRepository databaseports.OrderRepository
}

func NewOrderUseCase(httpClient httpports.HttpRequester, orderRepository databaseports.OrderRepository) *OrderUseCase {
	return &OrderUseCase{
		httpClient,
		orderRepository,
	}
}

func (ou *OrderUseCase) Create(order *domain.Order, token string) (*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userId, err := shared.ExtractIDFromToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token or extract user ID: %w", err)
	}

	order.CustomerID = userId

	var productIDs []string
	for _, item := range order.Items {
		productIDs = append(productIDs, item.ProductID)
	}

	var products []*domain.Product
	getProductURL := fmt.Sprintf("%s?ids=%s", os.Getenv("CHECK_AVAILABILITY"), strings.Join(productIDs, ","))
	if err := ou.httpClient.Get(ctx, getProductURL, &products); err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	productMap := make(map[string]*domain.Product)
	for _, p := range products {
		productMap[p.ID] = p
	}

	var total float64
	for i, item := range order.Items {
		product, ok := productMap[item.ProductID]
		if !ok {
			return nil, fmt.Errorf("product not found: %s", item.ProductID)
		}
		if product.Stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product: %s", product.ID)
		}
		order.Items[i].Price = product.Price
		total += product.Price * float64(item.Quantity)
	}

	order.TotalAmount = total
	order.Status = "PENDING"
	order.IsPaid = false
	order.CreatedAt = time.Now()

	order, err = ou.orderRepository.Create(order)
	if err != nil {
		return nil, fmt.Errorf("error happened while saving the order: %w", err)
	}

	return order, nil
}

func (ou *OrderUseCase) UpdateOrder(orderID string, updatedOrder *domain.UpdateOrder) (*domain.Order, error) {
	order, err := ou.orderRepository.FindByID(orderID)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve order: %w", err)
	}

	if order == nil {
		return nil, domain.ErrOrderNotFound
	}

	if order.IsPaid {
		return nil, fmt.Errorf("cannot update paid order")
	}

	productIDs := []string{}
	if updatedOrder.Items != nil {

		for _, item := range updatedOrder.Items {
			productIDs = append(productIDs, item.ProductID)
		}

		var products []*domain.Product
		getProductURL := fmt.Sprintf("%s?ids=%s", os.Getenv("CHECK_AVAILABILITY"), strings.Join(productIDs, ","))
		if err := ou.httpClient.Get(context.Background(), getProductURL, &products); err != nil {
			return nil, fmt.Errorf("failed to fetch products: %w", err)
		}

		productMap := make(map[string]*domain.Product)
		for _, p := range products {
			productMap[p.ID] = p
		}

		var total float64
		for _, item := range updatedOrder.Items {
			product, ok := productMap[item.ProductID]
			if !ok {
				return nil, fmt.Errorf("product not found: %s", item.ProductID)
			}
			if product.Stock < item.Quantity {
				return nil, fmt.Errorf("insufficient stock for product: %s", product.ID)
			}
			total += float64(item.Quantity) * product.Price
		}

		order.Items = updatedOrder.Items
		order.TotalAmount = total
		order.ID = ""
	}

	if updatedOrder.Status != "" {
		if updatedOrder.Status == "shipped" && !order.IsPaid {
			return nil, fmt.Errorf("cannot set status to 'shipped' for unpaid order")
		}

		order.Status = updatedOrder.Status
		order.ID = ""
	}

	if updatedOrder.IsPaid {
		order.IsPaid = updatedOrder.IsPaid
		order.ID = ""
	}

	order.UpdatedAt = time.Now()

	updatedOrderResult, err := ou.orderRepository.Update(orderID, order)

	if err != nil {
		return nil, fmt.Errorf("failed to update order: %w", err)
	}

	updatedOrderResult.ID = orderID

	return updatedOrderResult, nil
}

func (ou *OrderUseCase) GetOrder(orderID string) (*domain.Order, error) {
	order, err := ou.orderRepository.FindByID(orderID)

	if err != nil {
		return nil, fmt.Errorf("failed to get order: %w", err)
	}

	if order == nil {
		return nil, domain.ErrOrderNotFound
	}

	return order, nil
}

func (ou *OrderUseCase) GetUserOrders(token string) ([]*domain.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	userId, err := shared.ExtractIDFromToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to validate token or extract user ID: %w", err)
	}

	orders, err := ou.orderRepository.FindUserOrder(userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get user orders: %w", err)
	}

	return orders, nil
}
