package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/leetcode-daily-rank/services/leetcode"
)

func GetUserProfile(c *gin.Context) {
	userSlug := c.Param("userSlug")

	u, err := leetcode.GetUserProfile(userSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": u,
		"msg":  "ok",
	})
}

func GetUsersByRank(c *gin.Context) {
	users, err := leetcode.GetUsersByRank()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": users,
		"msg":  "ok",
	})
}
