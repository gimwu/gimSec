package main

import (
	"gimSec/basic/logging"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {

		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		logging.Info("register in data :", name, telephone, password)

		if len(telephone) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code2": 422, "msg": "telephone must be 11"})
			return
		}
		if len(password) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code2": 422, "msg": "password must be > 6"})
			return
		}

		ctx.JSON(200, gin.H{
			"msg1": "login success",
		})
	})

	panic(r.Run())
}
