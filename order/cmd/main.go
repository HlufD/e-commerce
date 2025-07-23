package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/HlufD/order-ms/cmd/docs"
	adapters "github.com/HlufD/order-ms/internal/adapters/left/http"
	"github.com/HlufD/order-ms/internal/adapters/left/http/controllers"
	persistence "github.com/HlufD/order-ms/internal/adapters/right/persistence/mongo"
	"github.com/HlufD/order-ms/internal/core/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Order Service API
// @version 1.0
// @description API documentation for the Order service.
// @host localhost:4003
// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token (e.g., "Bearer eyJhbGciOi...")
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in the environment variables")
	}

	dbAdapter := persistence.NewDatabaseConnectionAdapter(mongoURI)
	client, err := dbAdapter.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	orderRepo := persistence.NewOrderRepositoryAdapter(client)
	httpClient := adapters.NewHttpClient("", time.Second*10)
	orderUseCase := usecases.NewOrderUseCase(httpClient, orderRepo)

	orderController := controllers.NewOrderController(*orderUseCase)

	r := chi.NewRouter()
	r.Mount("/api/v1/orders", orderController.Routes())

	// Swagger route - must be registered on the main router
	r.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	r.HandleFunc("GET /swagger/*", httpSwagger.WrapHandler)

	port := ":4003"
	fmt.Printf("Order service is running on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer func() {
		if err := dbAdapter.Close(); err != nil {
			log.Fatalf("Failed to close MongoDB connection: %v", err)
		}
		fmt.Println("MongoDB connection closed.")
	}()
}
