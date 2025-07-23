package domain

import (
	"time"
)

type Payment struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	OrderID     string    `json:"orderId" bson:"orderId"`
	Amount      float64   `json:"amount" bson:"amount"`
	Status      string    `json:"status" bson:"status"`
	Method      string    `json:"method" bson:"method"`
	Transaction string    `json:"transactionId" bson:"transactionId"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt"`
}
