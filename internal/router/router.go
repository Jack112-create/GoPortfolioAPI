package router

import (
	"github.com/Jack112-create/GoPortfolioAPI/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/profile", handlers.GetGithubProfile)
}
