package job

import (
	"svc-loker-be/internal/models"
	"svc-loker-be/internal/response"
	"svc-loker-be/internal/services/job"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	jobSvc job.Service
}

func NewHandler(jobSvc job.Service) *Handler {
	return &Handler{jobSvc: jobSvc}
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	var filters models.JobFilter
	if err := c.QueryParser(&filters); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request parameters")
	}

	result, err := h.jobSvc.GetList(filters)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, err.Error())
	}

	return response.Success(c, result)
}
