package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
)

func AddAdmin(c *gin.Context) {
	var admin *model.Admin
	utils.BindJson(c, &admin)

	isExist, err := server.CheckAdmin(admin)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
		return
	}

	if isExist {
		logging.Info("register Telephone is exist")
		c.JSON(500, gin.H{
			"code": 500,
			"data": "register Telephone is exit",
		})
		return
	}

	err = server.AddAdmin(admin)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": admin,
	})
}

func AdminLogin(c *gin.Context) {
	var admin model.Admin
	utils.BindJson(c, &admin)

	err := server.AdminLogin(&admin)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "500",
			"data": err,
			"msg":  "login fail ",
		})
		return
	}

	token, err := jwt.ReleaseToken(admin.StateFullEntity)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "500",
			"data": err,
			"msg":  "login fail ",
		})
		return
	}

	c.Header("Authorization", token)

	c.JSON(200, gin.H{
		"code":          200,
		"data":          admin,
		"msg":           "Login success",
		"Authorization": token,
	})
}
