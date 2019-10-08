package api

import "github.com/gin-gonic/gin"

// Ping : Status check
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"massage": "pong",
	})
}
