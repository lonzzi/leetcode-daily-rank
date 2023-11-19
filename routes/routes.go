package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/leetcode-daily-rank/handlers"
	"github.com/lonzzi/leetcode-daily-rank/middlewares"
)

func InitRoute(r *gin.Engine) {
	r.Use(middlewares.Cors())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userG := r.Group("/user")
	{
		userG.GET("/:userSlug", handlers.GetUserProfile)
		userG.GET("/rank", handlers.GetUsersByRank)
	}

	r.GET("login", handlers.Login)
}
