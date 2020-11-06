package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kalifun/gin-template/global"
	"github.com/kalifun/gin-template/middleware/cors"
	v1 "github.com/kalifun/gin-template/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(logs.LogMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CorsMiddleware())
	gin.SetMode(global.ConfigSvr.System.Mode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/test", v1.Hello)
	}
	return r
}
