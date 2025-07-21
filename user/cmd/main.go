package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	httpAdapter "github.com/HlufD/users-ms/internals/adapters/left/http"
	adapters "github.com/HlufD/users-ms/internals/adapters/right"
	"github.com/HlufD/users-ms/internals/adapters/right/persistence/postgres"
	"github.com/HlufD/users-ms/internals/application"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dataSource := os.Getenv("POSTGRES_DSN")
	driver := os.Getenv("DRIVER_NAME")
	secrete := os.Getenv("SECRET")

	databaseAdapter := postgres.NewDatabaseConnectionAdapter(dataSource, driver)

	db, err := databaseAdapter.Connect()

	if err != nil {
		log.Fatalf("Failed to connect to database err: %v", err)
	}

	userRepository := postgres.NewUserRepositoryAdapter(db)
	hashingAdapter := adapters.NewBcryptAdapter(10)
	jwtAdapter := adapters.NewJWTAdapter(secrete, time.Hour)

	authService := application.NewAuthService(userRepository, hashingAdapter, jwtAdapter)
	authController := httpAdapter.NewAuthHandler(*authService)

	router := http.NewServeMux()

	router.HandleFunc("POST /register", func(w http.ResponseWriter, r *http.Request) {
		authController.Register(w, r)
	})
	router.HandleFunc("POST /login", func(w http.ResponseWriter, r *http.Request) {
		authController.Login(w, r)
	})

	// create server
	fmt.Printf("The server us running on port :%v\n", os.Getenv("PORT"))

	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)

	if err != nil {
		log.Fatal("Failed to start the server")
	}

}
