package models

import (
	"time"
)

type EmploymentType struct {
	ID        int       `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (EmploymentType) TableName() string {
	return "employment_types"
}

type EmploymentTypeResponse struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}
