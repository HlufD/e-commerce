package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/HlufD/products-ms/cmd/docs"
	"github.com/HlufD/products-ms/internal/adapters/left/http/controllers"
	persistence "github.com/HlufD/products-ms/internal/adapters/right/persistence/mongo"
	"github.com/HlufD/products-ms/internal/core/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Product Microservice API
// @version 1.0
// @description API for managing products
// @host localhost:4001
// @BasePath /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env variables.")
	}

	mongoURI := os.Getenv("MONGODB_URI")

	dbAdapter := persistence.NewDatabaseConnectionAdapter(mongoURI)

	client, err := dbAdapter.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer dbAdapter.Close()

	productRepo := persistence.NewProductRepositoryAdapter(client)
	productService := usecases.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	router := chi.NewRouter()

	// Register routes
	productController.RegisterRoutes(router)

	// Swagger route - must be registered on the main router
	router.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	router.HandleFunc("GET /swagger/*", httpSwagger.WrapHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Printf("Server running on http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
