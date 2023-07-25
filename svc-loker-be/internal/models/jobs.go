package models

import "time"

type Job struct {
	ID                 string    `json:"id;primaryKey"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	PositionLevelID    string    `json:"position_level_id"`
	PositionLevelName  string    `json:"position_level_name"`
	EmploymentTypeID   string    `json:"employment-type_id"`
	EmploymentTypeName string    `json:"employment_type_name"`
	Status             bool      `json:"status"`
	CreatedAt          time.Time `json:"column:created_at"`
	UpdatedAt          time.Time `json:"column:updated_at"`
}

type JobResponse struct {
	ID                 string `json:"id;primaryKey"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	PositionLevelName  string `json:"position_level_name"`
	EmploymentTypeName string `json:"employment_type_name"`
}
