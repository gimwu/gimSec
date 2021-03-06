package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/utils"
	v1 "gimSec/src/provider-user/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(utils.Cors())
	apiv1 := r.Group("/api/v1")
	apiv1.POST("/user", v1.AddUser)
	apiv1.POST("/user/login", v1.Login)

	apiv1.GET("/user/init", v1.InitUser)

	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.GET("/user/queryUserPage", v1.QueryUserPage)
		apiv1.PUT("/user", v1.EditUser)
		apiv1.GET("/user", v1.GetUser)
		apiv1.DELETE("/user", v1.DeleteUser)
		apiv1.PUT("/user/logout", v1.Logout)

		apiv1.POST("/user/addr", v1.AddAddr)
		apiv1.PUT("/user/addr", v1.EditAddr)
		apiv1.DELETE("/user/addr", v1.DeleteAddr)
		apiv1.GET("/user/queryMyAddr", v1.QueryMyAddr)
	}

	return r
}
