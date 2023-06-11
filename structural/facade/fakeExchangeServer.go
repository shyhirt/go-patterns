package facade

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func Run() {
	app := fiber.New()

	app.Get("/auth", func(c *fiber.Ctx) error {
		if c.Query("login") == "login" && c.Query("pass") == "pass" {
			return c.SendString("token")
		}
		c.Status(fiber.StatusForbidden)
		return c.SendString("unauthorized")
	})
	app.Get("/balance", func(c *fiber.Ctx) error {
		if c.Query("token") == "token" {
			return c.SendString("10000")
		}
		c.Status(fiber.StatusForbidden)
		return c.SendString("unauthorized")
	})
	// Other routes for test purposes
	// Add here buy and sell endpoints
	app.Get("/:id", func(c *fiber.Ctx) error {
		if c.Query("token") != "token" {
			c.Status(fiber.StatusForbidden)
			return c.SendString("unauthorized")
		}
		return c.SendString("10")
	})
	log.Fatal(app.Listen(":3000"))
}
