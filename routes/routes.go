package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/leetcode-daily-rank/handlers"
)

func InitRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userG := r.Group("/user")
	{
		userG.GET("/:userSlug", handlers.GetUserProfile)
	}
}
