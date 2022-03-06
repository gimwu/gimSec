package main

import (
	"gimSec/basic/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db, err := utils.GormMysqlConnection("localhost", 3306, "gimmick", "123456", "test")
	if err != nil {
		panic("error :" + err.Error())
	}
	log.Println("database connection success!", db)

	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {

		name := ctx.PostForm("name")
		telephone := ctx.PostForm("telephone")
		password := ctx.PostForm("password")

		log.Println("register in data :", name, telephone, password)

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
