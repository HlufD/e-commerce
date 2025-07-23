package domain

import "time"

type Order struct {
	ID            string      `json:"id" bson:"_id,omitempty"`
	CustomerID    string      `json:"customerId" bson:"customerId"`
	Status        string      `json:"status" bson:"status"`
	Items         []OrderItem `json:"items" bson:"items"`
	TotalAmount   float64     `json:"totalAmount" bson:"totalAmount"`
	PaymentMethod string      `json:"paymentMethod" bson:"paymentMethod"`
	IsPaid        bool        `json:"isPaid" bson:"isPaid"`
	CreatedAt     time.Time   `json:"createdAt" bson:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt" bson:"updatedAt"`
}

type UpdateOrder struct {
	Status        string      `json:"status,omitempty" bson:"status,omitempty"`
	Items         []OrderItem `json:"items,omitempty" bson:"items,omitempty"`
	TotalAmount   float64     `json:"totalAmount,omitempty" bson:"totalAmount,omitempty"`
	PaymentMethod string      `json:"paymentMethod,omitempty" bson:"paymentMethod,omitempty"`
	IsPaid        bool        `json:"isPaid,omitempty" bson:"isPaid,omitempty"`
}
