package server

import (
	_ "fitness-app/internal/docs"

	"fitness-app/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Run starts the Gin server and registers all routes
func Run() error {
	r := gin.Default()

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// API v1 group
	api := r.Group("/api/v1")
	{
		api.POST("/users", handlers.CreateUser)
		api.GET("/users/:id", handlers.GetUser)

		api.POST("/workouts", handlers.CreateWorkout)
		api.GET("/workouts/:id", handlers.GetWorkout)

		api.POST("/sessions", handlers.CreateSession)
		api.GET("/sessions/:id", handlers.GetSession)
	}

	// Start server on port 8080
	return r.Run(":8080")
}
