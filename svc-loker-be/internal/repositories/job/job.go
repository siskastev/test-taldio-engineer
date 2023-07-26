package job

import (
	"strings"
	"svc-loker-be/internal/models"

	"gorm.io/gorm"
)

type jobRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &jobRepository{db: db}
}

func (j *jobRepository) GetList(filters models.JobFilter) ([]models.Job, error) {
	var jobs []models.Job

	query := j.db.Model(&models.Job{}).
		Joins("JOIN position_levels ON jobs.position_level_id = position_levels.id").
		Joins("JOIN employment_types ON jobs.employment_type_id = employment_types.id")

	if filters.Title != "" {
		query = query.Where("jobs.title LIKE ?", "%"+filters.Title+"%")
	}

	if filters.PositionLevelID != "" {
		positionLevelIDs := ConvertStringToStringSlice(filters.PositionLevelID)
		query = query.Where("jobs.position_level_id IN (?)", positionLevelIDs)
	}

	if filters.EmploymentTypeID != "" {
		employmentTypeIDs := ConvertStringToStringSlice(filters.EmploymentTypeID)
		query = query.Where("jobs.employment_type_id IN (?)", employmentTypeIDs)
	}

	query = query.Where("status = ?", true)

	if err := query.Select([]string{
		"jobs.id",
		"jobs.title",
		"jobs.description",
		"position_levels.position_level",
		"employment_types.type",
	}).Find(&jobs).Error; err != nil {
		return nil, err
	}

	return jobs, nil
}

// ConvertStringToStringSlice converts a comma-separated string to a string slice
func ConvertStringToStringSlice(s string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, ",")
}
