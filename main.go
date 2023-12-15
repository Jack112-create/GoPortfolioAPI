package main

import (
	"github.com/gofiber/fiber/v2"
)

type GithubProfile struct {
	Login string `json:"login"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		// request := fiber.Get("https://api.github.com/users/Jack112-create")
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
