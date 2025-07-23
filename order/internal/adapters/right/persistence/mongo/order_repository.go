package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/HlufD/order-ms/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepositoryAdapter struct {
	db *mongo.Client
}

func NewOrderRepositoryAdapter(db *mongo.Client) *OrderRepositoryAdapter {
	return &OrderRepositoryAdapter{
		db,
	}
}

func (ora *OrderRepositoryAdapter) collection() *mongo.Collection {
	return ora.db.Database("order-ms").Collection("orders")
}

func (ora *OrderRepositoryAdapter) Create(order *domain.Order) (*domain.Order, error) {
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	result, err := ora.collection().InsertOne(context.Background(), order)
	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		order.ID = oid.Hex()
	} else {
		return nil, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	return order, nil
}

func (ora *OrderRepositoryAdapter) FindByID(id string) (*domain.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var order domain.Order

	err = ora.collection().FindOne(context.Background(), bson.M{"_id": objID}).Decode(&order)

	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (ora *OrderRepositoryAdapter) FindUserOrder(userId string) ([]*domain.Order, error) {
	cursor, err := ora.collection().Find(context.Background(), bson.M{"customerId": userId})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var orders []*domain.Order

	for cursor.Next(context.Background()) {
		var order domain.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, err
		}
		orders = append(orders, &order)
	}

	return orders, nil
}

func (ora *OrderRepositoryAdapter) Update(id string, order *domain.Order) (*domain.Order, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	order.UpdatedAt = time.Now()

	update := bson.M{
		"$set": order,
	}

	_, err = ora.collection().UpdateByID(context.Background(), objID, update)
	if err != nil {
		return nil, err
	}

	return order, nil
}
