package main

import (
	"log"
	"os"

	"github.com/Jack112-create/GoPortfolioAPI/middleware"
	"github.com/Jack112-create/GoPortfolioAPI/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	log.Println("Starting Go application")

	// Initialize app
	app := fiber.New()
	log.Println("New Fiber app started")

	// Grab API key from .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	middleware.ApiKey = os.Getenv("KEY")

	// Key authentication
	// app.Use(keyauth.New(keyauth.Config{
	// 	KeyLookup: "cookie:access_token",
	// 	Validator: middleware.ValidateAPIKey,
	// }))

	// Setup routes
	router.SetupRoutes(app)
	log.Println("Router setup successfully")

	app.Listen(":3000")
}
