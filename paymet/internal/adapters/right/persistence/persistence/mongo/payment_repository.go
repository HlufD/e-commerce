package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/HlufD/payment-ms/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PaymentRepositoryAdapter struct {
	db *mongo.Client
}

func NewOrderRepositoryAdapter(db *mongo.Client) *PaymentRepositoryAdapter {
	return &PaymentRepositoryAdapter{
		db,
	}
}

func (ora *PaymentRepositoryAdapter) collection() *mongo.Collection {
	return ora.db.Database("payment-ms").Collection("payments")
}

func NewPaymentRepositoryAdapter(db *mongo.Client) *PaymentRepositoryAdapter {
	return &PaymentRepositoryAdapter{
		db,
	}
}

func (pra *PaymentRepositoryAdapter) Create(payment *domain.Payment) (*domain.Payment, error) {
	payment.CreatedAt = time.Now()
	payment.UpdatedAt = time.Now()

	result, err := pra.collection().InsertOne(context.Background(), payment)
	if err != nil {
		return nil, err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		payment.ID = id.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}

	return payment, nil
}

func (pra *PaymentRepositoryAdapter) FindByID(id string) (*domain.Payment, error) {
	var payment domain.Payment

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ObjectID: %v", err)
	}

	err = pra.collection().FindOne(context.Background(), bson.M{"_id": objID}).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, err
	}

	return &payment, nil
}

func (pra *PaymentRepositoryAdapter) FindByOrderID(orderID string) (*domain.Payment, error) {
	var payment domain.Payment

	objID, err := primitive.ObjectIDFromHex(orderID)
	if err != nil {
		return nil, fmt.Errorf("invalid order ID format: %w", err)
	}

	query := bson.M{"orderId": objID}

	err = pra.collection().FindOne(context.Background(), query).Decode(&payment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domain.ErrPaymentNotFound
		}
		return nil, fmt.Errorf("failed to fetch payment: %w", err)
	}

	return &payment, nil
}
