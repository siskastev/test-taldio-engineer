package models

import "time"

type PositionLevel struct {
	ID        int       `gorm:"column:id"`
	Level     string    `gorm:"column:position_level"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (PositionLevel) TableName() string {
	return "position_levels"
}

type PositionLevelResponse struct {
	ID    int    `json:"id"`
	Level string `json:"level"`
}
