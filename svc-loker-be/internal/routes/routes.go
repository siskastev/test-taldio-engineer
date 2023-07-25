package routes

import (
	employee_type "svc-loker-be/internal/routes/employment_type"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	employee_type.RegisterRoutes(api)
}
