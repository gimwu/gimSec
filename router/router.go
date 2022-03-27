package router

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	v1 "gimSec/router/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(Cors())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")

	apiv1.POST("/user", v1.AddUser)
	apiv1.POST("/user/login", v1.Login)
	apiv1.POST("/admin", v1.AddAdmin)
	apiv1.POST("/admin/login", v1.AdminLogin)
	apiv1.Use(jwt.AuthMiddleware())
	{
		apiv1.GET("/admin/queryAdminPage", v1.QueryAdminPage)
		apiv1.GET("/admin", v1.GetAdmin)
		apiv1.PUT("/admin", v1.EditAdmin)
		apiv1.DELETE("/admin", v1.DeleteAdmin)

		apiv1.PUT("/user", v1.EditUser)
		apiv1.GET("/user", v1.GetUser)
		apiv1.DELETE("/user", v1.DeleteUser)
		apiv1.GET("/user/queryUserPage", v1.QueryUserPage)

		apiv1.POST("/goods", v1.AddGoods)
		apiv1.DELETE("/goods", v1.DeleteGoods)
		apiv1.PUT("/goods", v1.EditGoods)
		apiv1.GET("/goods", v1.GetGoods)
		apiv1.GET("/queryGoodsPage", v1.QueryGoodsPage)

		apiv1.POST("/goodsOrderItem", v1.AddGoodsOrderItem)
		apiv1.DELETE("/goodsOrderItem", v1.DeleteGoodsOrderItem)
		apiv1.GET("/goodsOrderItem", v1.GetGoodsOrderItem)
		apiv1.GET("/queryGoodsOrderItem", v1.QueryGoodsOrderItem)

		apiv1.POST("/order", v1.AddOrder)
	}

	return r
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "*")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}

		defer func() {
			if err := recover(); err != nil {
				logging.Error("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
