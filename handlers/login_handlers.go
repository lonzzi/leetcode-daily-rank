package handlers

import (
	"regexp"

	"github.com/gin-gonic/gin"
)

// only for test handler
func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		c.JSON(200, gin.H{
			"code":    0,
			"data":    nil,
			"message": "username or password is empty",
		})
		return
	}

	ValidateEmail := func(email string) bool {
		re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
		return re.MatchString(email)
	}

	if !ValidateEmail(username) {
		c.JSON(200, gin.H{
			"code":    0,
			"data":    nil,
			"message": "username is not email",
		})
		return
	}

	c.JSON(200, gin.H{
		"code":    1,
		"data":    nil,
		"message": "ok",
	})
}
