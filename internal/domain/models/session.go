package models

import "time"

type Session struct {
	ID         uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	WorkoutID  uint      `json:"workout_id" example:"1"`
	StartedAt  time.Time `json:"started_at" example:"2025-11-20T10:00:00Z"`
	FinishedAt time.Time `json:"finished_at" example:"2025-11-20T11:00:00Z"`
	CreatedAt  time.Time `json:"created_at" example:"2025-11-20T10:00:00Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2025-11-20T10:00:00Z"`
}
