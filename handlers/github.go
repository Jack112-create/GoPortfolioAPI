package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Jack112-create/GoPortfolioAPI/model"
	"github.com/gofiber/fiber/v2"
)

var (
	GithubUser model.GithubProfile
	GithubRepo model.GithubRepository
)

func GetGithubProfile(c *fiber.Ctx) error {
	log.Println("Received request to /profile")
	agent := fiber.Get("https://api.github.com/users/Jack112-create")
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	err := json.Unmarshal(body, &GithubUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	return c.Status(statusCode).JSON(GithubUser)
}

func GetGithubRepo(c *fiber.Ctx) error {
	log.Println("Received request to /repo/:name")
	repo := c.Params("name")
	repoUrl := fmt.Sprintf("https://api.github.com/repos/Jack112-create/%v", repo)

	agent := fiber.Get(repoUrl)
	statusCode, body, errs := agent.Bytes()
	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errs": errs,
		})
	}

	err := json.Unmarshal(body, &GithubRepo)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err,
		})
	}

	if GithubRepo.URL == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": "Repository does not exist",
		})
	}

	return c.Status(statusCode).JSON(GithubRepo)
}
