package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/HlufD/users-ms/cmd/docs" // Make sure this path is correct
	httpAdapter "github.com/HlufD/users-ms/internals/adapters/left/http"
	adapters "github.com/HlufD/users-ms/internals/adapters/right"
	"github.com/HlufD/users-ms/internals/adapters/right/persistence/postgres"
	"github.com/HlufD/users-ms/internals/application"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Users Microservice API
// @version 1.0
// @description API for user authentication and management
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database setup
	dataSource := os.Getenv("POSTGRES_DSN")
	driver := os.Getenv("DRIVER_NAME")
	secrete := os.Getenv("SECRET")

	databaseAdapter := postgres.NewDatabaseConnectionAdapter(dataSource, driver)
	db, err := databaseAdapter.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database err: %v", err)
	}

	// Service setup
	userRepository := postgres.NewUserRepositoryAdapter(db)
	hashingAdapter := adapters.NewBcryptAdapter(10)
	jwtAdapter := adapters.NewJWTAdapter(secrete, time.Hour)
	authService := application.NewAuthService(userRepository, hashingAdapter, jwtAdapter)
	authController := httpAdapter.NewAuthHandler(*authService)

	// Router setup
	router := http.NewServeMux()

	// Register routes
	router.HandleFunc("POST /api/v1/register", authController.Register)
	router.HandleFunc("POST /api/v1/login", authController.Login)

	// Swagger route - must be registered on the main router
	router.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)
	router.HandleFunc("GET /swagger/*", httpSwagger.WrapHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Printf("The server is running on port :%v\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
