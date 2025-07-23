package domain

type OrderItem struct {
	ProductID string  `json:"productId" bson:"productId"`
	Quantity  int     `json:"quantity" bson:"quantity"`
	Price     float64 `json:"price" bson:"price"`
}
