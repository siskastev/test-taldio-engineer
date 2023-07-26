package position_level

import (
	"svc-loker-be/internal/models"

	"gorm.io/gorm"
)

type positionLevelRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &positionLevelRepository{db: db}
}

func (e *positionLevelRepository) GetList() ([]models.PositionLevel, error) {
	var positionLevel []models.PositionLevel
	if err := e.db.Select("id", "position_level").Find(&positionLevel).Error; err != nil {
		return nil, err
	}

	return positionLevel, nil
}
