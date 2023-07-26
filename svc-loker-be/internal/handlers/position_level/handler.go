package position_level

import (
	"svc-loker-be/internal/response"
	"svc-loker-be/internal/services/position_level"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	levelService position_level.Service
}

func NewHandler(levelService position_level.Service) *Handler {
	return &Handler{levelService: levelService}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	result, err := h.levelService.GetList()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}
