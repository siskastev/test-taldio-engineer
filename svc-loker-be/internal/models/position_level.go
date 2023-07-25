package models

import "time"

type PositionLevel struct {
	ID        int       `json:"id"`
	Level     string    `json:"position_level"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (PositionLevel) TableName() string {
	return "position_levels"
}

type PositionLevelResponse struct {
	ID    int    `json:"id"`
	Level string `json:"level"`
}
