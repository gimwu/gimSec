package router

import (
	"gimSec/basic/utils"
	v1 "gimSec/src/provider-goods/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(utils.Cors())

	apiv1.POST("/goods", v1.AddGoods)
	apiv1.DELETE("/goods", v1.DeleteGoods)
	apiv1.PUT("/goods", v1.EditGoods)
	apiv1.GET("/goods", v1.GetGoods)
	apiv1.GET("/goods/queryGoodsPage", v1.QueryGoodsPage)

	return r
}
