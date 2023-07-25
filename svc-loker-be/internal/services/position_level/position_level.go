package position_level

import (
	"svc-loker-be/internal/models"
	"svc-loker-be/internal/repositories/position_level"
)

type positionLevelService struct {
	positionLevel position_level.Repository
}

func NewService(positionLevel position_level.Repository) Service {
	return &positionLevelService{positionLevel: positionLevel}
}

func (p *positionLevelService) GetList() ([]models.PositionLevelResponse, error) {
	result, err := p.positionLevel.GetList()
	if err != nil {
		return nil, err
	}

	var response []models.PositionLevelResponse
	for _, v := range result {
		levelResponse := models.PositionLevelResponse{
			ID:    v.ID,
			Level: v.Level,
		}

		response = append(response, levelResponse)
	}

	return response, nil
}
