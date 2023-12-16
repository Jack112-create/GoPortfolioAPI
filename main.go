package main

import (
	"github.com/Jack112-create/GoPortfolioAPI/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize app
	app := fiber.New()

	// Setup routes
	router.SetupRoutes(app)

	app.Listen(":3000")
}
