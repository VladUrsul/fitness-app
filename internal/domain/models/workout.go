package models

import "time"

type Workout struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" example:"1"`
	Type      string    `json:"type" example:"Yoga"`
	Scheduled time.Time `json:"scheduled" example:"2025-11-20T10:00:00Z"`
	CreatedAt time.Time `json:"created_at" example:"2025-11-20T10:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2025-11-20T10:00:00Z"`
}
