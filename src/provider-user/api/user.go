package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/provider-user/model"
	"gimSec/src/provider-user/server"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func AddUser(c *gin.Context) {

	var user *model.User
	utils.BindJson(c, &user)

	isExist, err := server.CheckUser(user)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	user.LastTimeLogin = time.Now()
	if isExist {
		logging.Info("register Telephone is exist")
		response.Error(c, "register Telephone is exit")
		return
	}

	err = server.AddUser(user)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, user)
}

func Login(c *gin.Context) {
	var user model.User
	utils.BindJson(c, &user)

	err := server.Login(&user)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	token, err := jwt.ReleaseToken(user.StateFullEntity)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	c.Header("Authorization", token)

	response.Success(c, user, token)
}

func EditUser(c *gin.Context) {
	id := c.Query("id")
	logging.Debug(id)

	user, err := server.GetUser(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	utils.BindJson(c, &user)

	err = server.EditUser(user)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, user)

}

func GetUser(c *gin.Context) {
	id := c.Query("id")
	logging.Debug(id)

	user, err := server.GetUser(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	user, err := server.DeleteUser(id)
	if err != nil {
		logging.Error("DeleteUser error :", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, user)
}

func QueryUserPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	logging.Debug(currentPage, pageSize)
	data, err := server.QueryUserPage(nil, currentPage, pageSize)

	if err != nil {
		logging.Error("QueryUserPage Error:", err)
		response.Error(c, err.Error())
	}
	response.Success(c, data)
}
