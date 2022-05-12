package router

import (
	"gimSec/basic/utils"
	v1 "gimSec/src/consumer-sec-order/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(utils.Cors())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/secOrder", v1.AddSecOrder)
		apiv1.GET("/secOrder/querySecOrderPage", v1.QuerySecOrderPage)
		apiv1.GET("/secOrder/queryMyOrder", v1.QueryMyOrder)
	}
	return r
}
