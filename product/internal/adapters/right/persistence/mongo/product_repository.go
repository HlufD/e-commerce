package persistence

import (
	"context"
	"errors"
	"time"

	"github.com/HlufD/products-ms/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryAdapter struct {
	db *mongo.Client
}

func (pr *ProductRepositoryAdapter) collection() *mongo.Collection {
	return pr.db.Database("product-ms").Collection("products")
}

func NewProductRepositoryAdapter(db *mongo.Client) *ProductRepositoryAdapter {
	return &ProductRepositoryAdapter{
		db,
	}
}

func (pr *ProductRepositoryAdapter) Save(product *domain.Product) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	// add created and updated at
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := pr.collection().InsertOne(ctx, product)

	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		product.ID = oid.Hex()
	}

	return product, nil
}

func (pr *ProductRepositoryAdapter) Update(id string, updateProduct *domain.UpdateProduct) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": oid}
	update := bson.M{"$set": updateProduct}

	_, err = pr.collection().UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}

	//convert updateProduct to product
	product := &domain.Product{
		ID:          oid.Hex(),
		Name:        updateProduct.Name,
		Description: updateProduct.Description,
		Price:       updateProduct.Price,
		Stock:       updateProduct.Stock,
		Category:    updateProduct.Category,
		UpdatedAt:   time.Now(),
	}

	return product, nil

}

func (pr *ProductRepositoryAdapter) GetProductById(id string) (*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var product *domain.Product = &domain.Product{}

	err = pr.collection().FindOne(ctx, bson.M{"_id": oid}).Decode(product)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
		return nil, err
	}

	return product, nil
}

func (pr *ProductRepositoryAdapter) GetAllProducts() ([]*domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := pr.collection().Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var products []*domain.Product
	for cursor.Next(ctx) {
		var product domain.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pr *ProductRepositoryAdapter) CheckAvailability(id string, quantity int) (bool, error) {
	product, err := pr.GetProductById(id)
	if err != nil {
		return false, err
	}
	if product == nil {
		return false, nil
	}
	return product.Stock >= quantity, nil
}
