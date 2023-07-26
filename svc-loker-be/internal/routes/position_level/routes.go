package position_level

import (
	"svc-loker-be/internal/database"
	handlerPositionLevel "svc-loker-be/internal/handlers/position_level"
	repoPositionLevel "svc-loker-be/internal/repositories/position_level"
	servicePositionLevel "svc-loker-be/internal/services/position_level"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(route fiber.Router) {
	repo := repoPositionLevel.NewRepository(database.DB)
	service := servicePositionLevel.NewService(repo)
	handler := handlerPositionLevel.NewHandler(service)
	route.Get("/position-levels", handler.GetAll)
}
