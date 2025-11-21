package handlers

import (
	"fitness-app/internal/domain/models"
	"fitness-app/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateWorkout godoc
// @Summary Create a workout
// @Tags Workouts
// @Accept json
// @Produce json
// @Param workout body models.Workout true "Workout"
// @Success 200 {object} models.Workout
// @Router /api/v1/workouts [post]
func CreateWorkout(c *gin.Context) {
	var w models.Workout
	if err := c.ShouldBindJSON(&w); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	_, err := services.GetUser(w.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	// Create the workout
	created, err := services.CreateWorkout(&w)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create workout"})
		return
	}

	c.JSON(http.StatusOK, created)
}

// GetWorkout godoc
// @Summary Get a workout by ID
// @Tags Workouts
// @Produce json
// @Param id path int true "Workout ID"
// @Success 200 {object} models.Workout
// @Failure 404 {object} map[string]string
// @Router /api/v1/workouts/{id} [get]
func GetWorkout(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid workout id"})
		return
	}

	w, err := services.GetWorkout(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "workout not found"})
		return
	}

	c.JSON(http.StatusOK, w)
}
