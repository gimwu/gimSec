package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
	"strconv"
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

func Login(c *gin.Context) {
	var user model.User
	utils.BindJson(c, &user)

	err := server.Login(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"code": "500",
			"data": err,
			"msg":  "login fail ",
		})
		return
	}

	token, err := jwt.ReleaseToken(user)
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
		"data":          user,
		"msg":           "Login success",
		"Authorization": token,
	})
}

func EditUser(c *gin.Context) {
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

	utils.BindJson(c, &user)

	err = server.EditUser(user)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": "500",
			"data": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": "200",
		"data": user,
		"msg":  "success",
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

func QueryUserPage(c *gin.Context) {
	json := make(map[string]interface{})
	utils.BindJson(c, &json)
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	logging.Debug(currentPage, pageSize)
	data, err := server.QueryUserPage(&json, currentPage, pageSize)

	if err != nil {
		logging.Error("QueryUserPage Error:", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
