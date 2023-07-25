package employment_type

import "svc-loker-be/internal/models"

type Repository interface {
	GetList() ([]models.EmploymentType, error)
}
