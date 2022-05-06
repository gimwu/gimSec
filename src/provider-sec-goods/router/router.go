package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(utils.Cors())
	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.POST("/secGoods", api.AddSecGoods)
	}
	return r
}