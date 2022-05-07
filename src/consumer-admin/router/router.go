package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/utils"
	v1 "gimSec/src/consumer-admin/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(utils.Cors())

	apiv1 := r.Group("/api/v1")

	apiv1.POST("/admin", v1.AddAdmin)
	apiv1.POST("/admin/login", v1.AdminLogin)
	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.GET("/admin/queryAdminPage", v1.QueryAdminPage)
		apiv1.GET("/admin", v1.GetAdmin)
		apiv1.PUT("/admin", v1.EditAdmin)
		apiv1.DELETE("/admin", v1.DeleteAdmin)
	}

	return r
}
