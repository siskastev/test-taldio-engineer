package routes

import (
	"svc-loker-be/internal/routes/employment_type"
	"svc-loker-be/internal/routes/job"
	"svc-loker-be/internal/routes/position_level"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	employment_type.RegisterRoutes(api)
	position_level.RegisterRoutes(api)
	job.RegisterRoutes(api)
}
