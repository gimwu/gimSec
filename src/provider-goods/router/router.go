package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/utils"
	v1 "gimSec/src/provider-goods/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(utils.Cors())
	apiv1 := r.Group("/api/v1")

	apiv1.GET("/goods/queryGoodsPage", v1.QueryGoodsPage)
	apiv1.GET("/goods", v1.GetGoods)
	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.POST("/goods", v1.AddGoods)
		apiv1.DELETE("/goods", v1.DeleteGoods)
		apiv1.PUT("/goods", v1.EditGoods)

		apiv1.GET("/goods/getMyCart", v1.GetMyCart)
		apiv1.POST("/goods/addCart", v1.AddCart)
		apiv1.DELETE("/godos/deleteCart", v1.DeleteCart)
	}

	return r
}
