package services

import (
	"errors"
	"fitness-app/internal/db"
	"fitness-app/internal/domain/models"
)

func CreateWorkout(w *models.Workout) (*models.Workout, error) {
	var user models.User
	if err := db.DB.First(&user, w.UserID).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := db.DB.Create(w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

func GetWorkout(id uint) (*models.Workout, error) {
	var w models.Workout
	if err := db.DB.First(&w, id).Error; err != nil {
		return nil, err
	}
	return &w, nil
}
