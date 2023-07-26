package job

import "svc-loker-be/internal/models"

type Repository interface {
	GetList(filters models.JobFilter) ([]models.Job, error)
}
