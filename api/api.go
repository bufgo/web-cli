package api

import (
	"bufgo/logger"
	"bufgo/serializer"
	"github.com/gin-gonic/gin"
)

// Ping : Status check
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg: "pong",
	})
	logger.Infof("Check health")
}
