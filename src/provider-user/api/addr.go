package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/provider-user/model"
	"gimSec/src/provider-user/server"
	"github.com/gin-gonic/gin"
)

func AddAddr(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	userId := claims.UserId

	var addr *model.Addr
	err = utils.BindJson(c, &addr)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	addr.UserId = userId
	err = server.AddAddr(addr)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, addr)
}

func EditAddr(c *gin.Context) {
	id := c.Query("id")

	addr, err := server.GetAddr(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	err = utils.BindJson(c, &addr)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	err = server.EditAddr(addr)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, addr)
}

func DeleteAddr(c *gin.Context) {
	var params map[string]string
	err := utils.BindJson(c, &params)
	if err != nil {
		response.Error(c, err.Error())
		logging.Error(err)
		return
	}
	user, err := server.DeleteAddr(params["id"])
	if err != nil {
		logging.Error("DeleteUser error :", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, user)
}

func QueryMyAddr(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	userId := claims.UserId
	params := make(map[string]string, 0)
	params["userId"] = userId
	data, err := server.QueryAddrList(params)

	if err != nil {
		logging.Error("QueryUserPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data)
}
