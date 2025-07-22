package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/HlufD/products-ms/internal/adapters/left/http/controllers"
	persistence "github.com/HlufD/products-ms/internal/adapters/right/persistence/mongo"
	"github.com/HlufD/products-ms/internal/core/usecases"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

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

	// Middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	// Register routes
	productController.RegisterRoutes(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	log.Printf("Server running on http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}

}
