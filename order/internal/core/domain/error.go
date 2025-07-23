package domain

import "errors"

var (
	// ===== Validation Errors =====
	ErrMissingOrderID       = errors.New("order ID is required")
	ErrInvalidOrderID       = errors.New("invalid order ID format")
	ErrMissingCustomerID    = errors.New("customer ID is required")
	ErrInvalidOrderStatus   = errors.New("invalid order status")
	ErrInvalidPaymentMethod = errors.New("invalid payment method")
	ErrMissingItems         = errors.New("order must contain at least one item")

	// ===== Order Errors =====
	ErrOrderNotFound       = errors.New("order not found")
	ErrDuplicateOrder      = errors.New("duplicate order")
	ErrOrderAlreadyPaid    = errors.New("order is already paid")
	ErrOrderAlreadyShipped = errors.New("order has already been shipped")
	ErrOrderCancelled      = errors.New("order has been cancelled")

	// ===== Inventory/Product Related Errors =====
	ErrProductNotFound   = errors.New("product in order not found")
	ErrOutOfStock        = errors.New("product is out of stock")
	ErrInsufficientStock = errors.New("insufficient stock for ordered item")
	ErrNegativeQuantity  = errors.New("ordered quantity cannot be negative")

	// ===== Payment Errors =====
	ErrPaymentFailed  = errors.New("payment processing failed")
	ErrInvalidPayment = errors.New("invalid payment details")
)
