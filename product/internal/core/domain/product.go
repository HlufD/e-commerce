package domain

import (
	"time"
)

type Product struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string    `bson:"name,omitempty" json:"name,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64   `bson:"price,omitempty" json:"price,omitempty"`
	Stock       int       `bson:"stock,omitempty" json:"stock,omitempty"`
	Category    string    `bson:"category,omitempty" json:"category,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}

type UpdateProduct struct {
	Name        string    `bson:"name,omitempty" json:"name,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Price       float64   `bson:"price,omitempty" json:"price,omitempty"`
	Stock       int       `bson:"stock,omitempty" json:"stock,omitempty"`
	Category    string    `bson:"category,omitempty" json:"category,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}
