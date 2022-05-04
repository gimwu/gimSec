package router

import (
	"gimSec/basic/utils"
	v1 "gimSec/src/provider-user/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(utils.Cors())
	apiv1.POST("/user", v1.AddUser)
	apiv1.POST("/user/login", v1.Login)

	apiv1.GET("/user/queryUserPage", v1.QueryUserPage)
	apiv1.PUT("/user", v1.EditUser)
	apiv1.GET("/user", v1.GetUser)
	apiv1.DELETE("/user", v1.DeleteUser)

	return r
}
