package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/ritesh-15/mapup-assessment/routes"
)

func main() {
	app := fiber.New()

	routes.Register(app)

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":  http.StatusOK,
			"message": "Health check successful!",
			"ok":      true,
		})
	})

	log.Fatal(app.Listen("localhost:8080"))
}
