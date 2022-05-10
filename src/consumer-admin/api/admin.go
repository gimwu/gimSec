package v1

import (
	"context"
	"gimSec/basic/global"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/consumer-admin/model"
	"gimSec/src/consumer-admin/server"
	"gimSec/src/provider-user/jwt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func AddAdmin(c *gin.Context) {
	var admin *model.Admin
	utils.BindJson(c, &admin)

	isExist, err := server.CheckAdmin(admin)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	if isExist {
		logging.Info("register Username is exist")
		response.Error(c, "register Username is exit")
		return
	}

	err = server.AddAdmin(admin)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, admin, nil)
}

func AdminLogin(c *gin.Context) {
	var admin model.Admin
	utils.BindJson(c, &admin)

	err := server.AdminLogin(&admin)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	token, err := jwt.ReleaseToken(admin.StateFullEntity)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	result, err := global.REDIS.Set(context.Background(), admin.Id, token, redis.KeepTTL).Result()
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	logging.Debug(result, "Success")

	c.Header("Authorization", token)

	response.Success(c, admin, token)
}

func QueryAdminPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	data, err := server.QueryAdminPage(nil, currentPage, pageSize)
	if err != nil {
		logging.Error("QueryAdminPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data)
}

func GetAdmin(c *gin.Context) {
	id := c.Query("id")

	admin, err := server.GetAdmin(id)

	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, admin)

}

func EditAdmin(c *gin.Context) {
	id := c.Query("id")

	admin, err := server.GetAdmin(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	utils.BindJson(c, &admin)

	err = server.EditAdmin(admin)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, admin)
}

func DeleteAdmin(c *gin.Context) {
	id := c.Query("id")
	admin, err := server.DeleteAdmin(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
	}

	response.Success(c, admin, nil)
}
