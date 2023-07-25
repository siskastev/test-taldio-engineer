package employment_type

import (
	"svc-loker-be/internal/models"
	"svc-loker-be/internal/repositories/employment_type"
)

type employmentTypeService struct {
	employmentTypeRepo employment_type.Repository
}

func NewService(employmentTypeRepo employment_type.Repository) Service {
	return &employmentTypeService{employmentTypeRepo: employmentTypeRepo}
}

func (p *employmentTypeService) GetList() ([]models.EmploymentTypeResponse, error) {
	result, err := p.employmentTypeRepo.GetList()
	if err != nil {
		return nil, err
	}

	var response []models.EmploymentTypeResponse
	for _, v := range result {
		typeResponse := models.EmploymentTypeResponse{
			ID:   v.ID,
			Type: v.Type,
		}

		response = append(response, typeResponse)
	}

	return response, nil
}
