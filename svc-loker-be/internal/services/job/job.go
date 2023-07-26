package job

import (
	"svc-loker-be/internal/models"
	"svc-loker-be/internal/repositories/job"
)

type jobService struct {
	jobRepo job.Repository
}

func NewService(jobRepo job.Repository) Service {
	return &jobService{jobRepo: jobRepo}
}

func (j *jobService) GetList(filters models.JobFilter) ([]models.JobResponse, error) {

	result, err := j.jobRepo.GetList(models.JobFilter{
		Title:            filters.Title,
		PositionLevelID:  filters.PositionLevelID,
		EmploymentTypeID: filters.EmploymentTypeID,
	})

	if err != nil {
		return nil, err
	}

	var response []models.JobResponse
	for _, v := range result {
		typeResponse := models.JobResponse{
			ID:                 v.ID,
			Title:              v.Title,
			Description:        v.Description,
			PositionLevelName:  v.PositionLevelName,
			EmploymentTypeName: v.EmploymentTypeName,
		}

		response = append(response, typeResponse)
	}

	return response, nil
}

func (j *jobService) GetByID(id string) (models.JobResponse, error) {
	result, err := j.jobRepo.GetByID(id)
	if err != nil {
		return models.JobResponse{}, err
	}

	return models.JobResponse{
		ID:                 result.ID,
		Title:              result.Title,
		Description:        result.Description,
		EmploymentTypeName: result.EmploymentTypeName,
		PositionLevelName:  result.PositionLevelName,
	}, err
}
