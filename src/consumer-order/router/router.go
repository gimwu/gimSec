package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/utils"
	v1 "gimSec/src/consumer-order/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(utils.Cors())

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.POST("/order", v1.AddOrder)
		apiv1.GET("/order/queryOrderPage", v1.QueryOrderPage)
		apiv1.GET("/order/queryMyOrder", v1.QueryMyOrder)
	}

	return r
}
