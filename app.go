package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sample GoLang Application")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("{ \"status\" : \"up\" }")
	})

	app.Listen("0.0.0.0:3000")
}
