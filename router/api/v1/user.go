package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {

	var user *model.User
	utils.BindJson(c, &user)

	isExist, err := server.Check(user)
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

	err = server.AddUser(user)
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
		"data": user,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	logging.Debug(id)

	user, err := server.GetUser(id)
	if err != nil {
		logging.Error(err)
		c.JSON(200, gin.H{
			"msg":  "500",
			"data": err,
		})
		return
	}

	token, err := jwt.ReleaseToken(*user)
	if err != nil {
		logging.Error("ReleaseToken error:", err)
		c.JSON(500, gin.H{
			"msg":  500,
			"data": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":   "200",
		"data":  user,
		"token": token,
	})
}
