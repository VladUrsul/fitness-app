package services

import (
	"errors"
	"fitness-app/internal/db"
	"fitness-app/internal/domain/models"
)

func CreateSession(s *models.Session) (*models.Session, error) {
	var w models.Workout
	if err := db.DB.First(&w, s.WorkoutID).Error; err != nil {
		return nil, errors.New("workout not found")
	}

	if err := db.DB.Create(s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func GetSession(id uint) (*models.Session, error) {
	var s models.Session
	if err := db.DB.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}
