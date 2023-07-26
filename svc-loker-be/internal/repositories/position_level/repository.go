package position_level

import "svc-loker-be/internal/models"

type Repository interface {
	GetList() ([]models.PositionLevel, error)
}
