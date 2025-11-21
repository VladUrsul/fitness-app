package handlers

import (
	"fitness-app/internal/domain/models"
	"fitness-app/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSession godoc
// @Summary Create a session
// @Tags Sessions
// @Accept json
// @Produce json
// @Param session body models.Session true "Session"
// @Success 200 {object} models.Session
// @Router /api/v1/sessions [post]
func CreateSession(c *gin.Context) {
	var s models.Session
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := services.CreateSession(&s)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, created)
}

// GetSession godoc
// @Summary Get a session by ID
// @Tags Sessions
// @Produce json
// @Param id path int true "Session ID"
// @Success 200 {object} models.Session
// @Failure 404 {object} map[string]string
// @Router /api/v1/sessions/{id} [get]
func GetSession(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session id"})
		return
	}

	s, err := services.GetSession(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "session not found"})
		return
	}

	c.JSON(http.StatusOK, s)
}
