package employment_type

import (
	"svc-loker-be/internal/models"

	"gorm.io/gorm"
)

type employmentTypeRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &employmentTypeRepository{db: db}
}

func (e *employmentTypeRepository) GetList() ([]models.EmploymentType, error) {
	var employmentType []models.EmploymentType
	if err := e.db.Select("id", "type").Find(&employmentType).Error; err != nil {
		return nil, err
	}

	return employmentType, nil
}
