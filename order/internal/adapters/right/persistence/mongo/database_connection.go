package persistence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnectionAdapter struct {
	client *mongo.Client
	uri    string
}

func NewDatabaseConnectionAdapter(uri string) *DatabaseConnectionAdapter {
	return &DatabaseConnectionAdapter{uri: uri}
}

func (db *DatabaseConnectionAdapter) Connect() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(db.uri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db.client = client

	return client, nil
}

func (db *DatabaseConnectionAdapter) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if db.client != nil {
		return db.client.Disconnect(ctx)
	}

	return nil
}
