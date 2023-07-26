package models

import (
	"time"
)

type EmploymentType struct {
	ID        int       `gorm:"column:id"`
	Type      string    `gorm:"column:type"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (EmploymentType) TableName() string {
	return "employment_types"
}

type EmploymentTypeResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}
