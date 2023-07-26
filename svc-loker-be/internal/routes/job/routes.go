package job

import (
	"svc-loker-be/internal/database"
	handlerJob "svc-loker-be/internal/handlers/job"
	repoJob "svc-loker-be/internal/repositories/job"
	serviceJob "svc-loker-be/internal/services/job"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(route fiber.Router) {
	repo := repoJob.NewRepository(database.DB)
	service := serviceJob.NewService(repo)
	handler := handlerJob.NewHandler(service)
	route.Get("/jobs", handler.GetAll)
}
