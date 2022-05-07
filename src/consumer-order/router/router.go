package router

import (
	"gimSec/basic/utils"
	v1 "gimSec/src/consumer-order/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(utils.Cors())

	apiv1.POST("/order", v1.AddOrder)

	return r
}
