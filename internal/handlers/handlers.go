package handlers

import (
	"encoding/json"

	"github.com/Jack112-create/GoPortfolioAPI/profile"
	"github.com/gofiber/fiber/v2"
)

func GetGithubProfile(c *fiber.Ctx) error {
	agent := fiber.Get("https://api.github.com/users/Jack112-create")
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	err := json.Unmarshal(body, &profile.GithubUser)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	return c.Status(statusCode).JSON(profile.GithubUser)
}
