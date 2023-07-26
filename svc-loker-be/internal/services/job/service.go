package job

import "svc-loker-be/internal/models"

type Service interface {
	GetList(filters models.JobFilter) ([]models.JobResponse, error)
	GetByID(id string) (models.JobResponse, error)
}
