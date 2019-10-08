package server

import (
	"bufgo/api"

	"github.com/gin-gonic/gin"
)

// NewRouter : Router setting
func NewRouter() *gin.Engine {
	router := gin.Default()

	// Router
	v1 := router.Group("api/v1")
	{
		v1.POST("ping", api.Ping)
	}
	return router
}
