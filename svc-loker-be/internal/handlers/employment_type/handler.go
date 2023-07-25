package employment_type

import (
	"svc-loker-be/internal/response"
	"svc-loker-be/internal/services/employment_type"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	typeService employment_type.Service
}

func NewHandler(typeService employment_type.Service) *Handler {
	return &Handler{typeService: typeService}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	result, err := h.typeService.GetList()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}
