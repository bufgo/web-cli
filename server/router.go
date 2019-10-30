package server

import (
	"bufgo/api"

	"github.com/gin-gonic/gin"
)

// NewRouter : 设置路由
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Router
	v1 := r.Group("api/v1")
	{
		v1.POST("ping", api.Ping)
	}
	return r
}
