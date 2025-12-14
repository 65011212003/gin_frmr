package router

import (
	"gin_frmr/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Service is healthy",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}
