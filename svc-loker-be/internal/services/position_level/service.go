package position_level

import "svc-loker-be/internal/models"

type Service interface {
	GetList() ([]models.PositionLevelResponse, error)
}
