package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/HlufD/payment-ms/cmd/docs"
	controller "github.com/HlufD/payment-ms/internal/adapters/left/http/controllers"
	persistence "github.com/HlufD/payment-ms/internal/adapters/right/persistence/persistence/mongo"
	"github.com/HlufD/payment-ms/internal/core/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Payment Service API
// @version 1.0
// @description This is a payment service for processing payments
// @host localhost:4004
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatal("MONGODB_URI is not set in the environment variables")
	}

	dbAdapter := persistence.NewDatabaseConnectionAdapter(mongoURI)
	client, err := dbAdapter.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	paymentRepo := persistence.NewPaymentRepositoryAdapter(client)
	paymentUseCase := usecases.NewPaymentUseCase(paymentRepo)

	paymentController := controller.NewPaymentController(paymentUseCase)

	r := chi.NewRouter()
	r.Mount("/api/v1/payments", paymentController.Routes(r))

	// Swagger route - must be registered on the main router
	r.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	r.HandleFunc("GET /swagger/*", httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4004"
	}
	fmt.Printf("Payment service is running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer func() {
		if err := dbAdapter.Close(); err != nil {
			log.Fatalf("Failed to close MongoDB connection: %v", err)
		}
		fmt.Println("MongoDB connection closed.")
	}()
}
