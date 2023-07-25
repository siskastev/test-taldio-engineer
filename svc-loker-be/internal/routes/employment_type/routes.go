package employment_type

import (
	"svc-loker-be/internal/database"
	handlerEmploymentType "svc-loker-be/internal/handlers/employment_type"
	repoEmploymentType "svc-loker-be/internal/repositories/employment_type"
	serviceEmploymentType "svc-loker-be/internal/services/employment_type"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(route fiber.Router) {
	repo := repoEmploymentType.NewRepository(database.DB)
	service := serviceEmploymentType.NewService(repo)
	handler := handlerEmploymentType.NewHandler(service)
	route.Get("/employment-types", handler.GetAll)
}
