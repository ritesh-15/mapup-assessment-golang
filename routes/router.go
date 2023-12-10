package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ritesh-15/mapup-assessment/controllers"
)

func Register(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/process-single", controllers.ProcessSingle)

	api.Post("/process-concurrent", controllers.ProcessConcurrent)
}
