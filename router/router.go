package router

import (
	"fmt"
	"gimSec/basic/jwt"
	v1 "gimSec/router/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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
		apiv1.GET("/admin/:id", v1.GetAdmin)
		apiv1.PUT("/admin/:id", v1.EditAdmin)
		apiv1.DELETE("/admin/:id", v1.DeleteAdmin)

		apiv1.PUT("/user", v1.EditUser)
		apiv1.GET("/user/:id", v1.GetUser)
		apiv1.DELETE("/user/:id", v1.DeleteUser)
		apiv1.GET("/user/queryUserPage", v1.QueryUserPage)

		apiv1.POST("/goods", v1.AddGoods)
		apiv1.DELETE("/goods/:id", v1.DeleteGoods)
		apiv1.PUT("/goods/:id", v1.EditGoods)
		apiv1.GET("/goods/:id", v1.GetGoods)
		apiv1.GET("/queryGoodsPage", v1.QueryGoodsPage)

		apiv1.POST("/goodsOrderItem", v1.AddGoodsOrderItem)
	}

	return r
}
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {

			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE,UPDATE") //服务器支持访的所有的跨域请求
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()

	}
}
