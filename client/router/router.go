package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/huiming23344/nanoservice/client/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	apiv1 := r.Group("/api")
	apiv1.Use()
	apiv1.GET("getInfo", v1.QueryInfo)
	return r
}
