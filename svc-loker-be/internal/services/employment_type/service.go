package employment_type

import "svc-loker-be/internal/models"

type Service interface {
	GetList() ([]models.EmploymentTypeResponse, error)
}
