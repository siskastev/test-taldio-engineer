package models

import "time"

type Job struct {
	ID                 string    `gorm:"column:id"`
	Title              string    `gorm:"column:title"`
	Description        string    `gorm:"column:description"`
	PositionLevelID    string    `gorm:"column:position_level_id"`
	PositionLevelName  string    `gorm:"column:position_level"`
	EmploymentTypeID   string    `gorm:"column:employment_type_id"`
	EmploymentTypeName string    `gorm:"column:type"`
	Status             bool      `gorm:"column:status"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}

type JobResponse struct {
	ID                 string `json:"id"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	PositionLevelName  string `json:"position_level_name"`
	EmploymentTypeName string `json:"employment_type_name"`
}

type JobFilter struct {
	Title            string `query:"title"`
	PositionLevelID  string `query:"level_id"`
	EmploymentTypeID string `query:"type_id"`
}
