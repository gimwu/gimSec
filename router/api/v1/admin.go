package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
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
		logging.Info("register Telephone is exist")
		response.Error(c, "register Telephone is exit")
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
		response.Error(c, err.Error())
		return
	}

	token, err := jwt.ReleaseToken(admin.StateFullEntity)
	if err != nil {
		response.Error(c, err.Error())
		return
	}

	c.Header("Authorization", token)

	c.JSON(200, gin.H{
		"code":          200,
		"data":          admin,
		"msg":           "Login success",
		"Authorization": token,
	})
	response.Success(c, admin, nil)
}

func QueryAdminPage(c *gin.Context) {
	json := make(map[string]interface{})
	utils.BindJson(c, &json)
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	data, err := server.QueryAdminPage(&json, currentPage, pageSize)
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
