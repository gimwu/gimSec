package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(utils.Cors())

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/secGoods/querySecGoodsPage", api.QuerySecGoodsPage)
	apiv1.GET("/secGoods", api.GetSecGoods)
	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.POST("/secGoods", api.AddSecGoods)
		apiv1.PUT("/secGoods", api.EditSecGoods)
		apiv1.DELETE("/secGoods", api.DeleteSecGoods)
	}
	return r
}
