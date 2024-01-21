package router

import (
	"github.com/Jack112-create/GoPortfolioAPI/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/profile", handlers.GetGithubProfile)
	app.Get("/repo/:name", handlers.GetGithubRepo)
}
