package domain

import "errors"

var (
	// Product errors
	ErrProductNotFound     = errors.New("product not found")
	ErrDuplicateProduct    = errors.New("product already exists (duplicate SKU/barcode)")
	ErrInvalidProductID    = errors.New("invalid product ID format")
	ErrOutOfStock          = errors.New("product out of stock")
	ErrInsufficientStock   = errors.New("insufficient stock for operation")
	ErrInvalidPrice        = errors.New("invalid product price")
	ErrInvalidCategory     = errors.New("invalid product category")
	ErrProductNotPublished = errors.New("product is not published/visible")
	ErrProductDeleted      = errors.New("product has been deleted")

	// Inventory errors
	ErrInventoryUpdateFailed = errors.New("failed to update inventory")
	ErrNegativeStock         = errors.New("stock cannot be negative")

	// Variant errors
	ErrVariantConflict      = errors.New("variant combination already exists")
	ErrInvalidVariantOption = errors.New("invalid variant option (e.g., size/color)")

	// Validation errors
	ErrMissingSKU         = errors.New("SKU/barcode is required")
	ErrMissingTitle       = errors.New("product title is required")
	ErrInvalidDescription = errors.New("invalid product description format")
)
