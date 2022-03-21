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

func QueryAdminPage(c *gin.Context) {
	json := make(map[string]interface{})
	utils.BindJson(c, &json)
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	data, err := server.QueryAdminPage(&json, currentPage, pageSize)
	if err != nil {
		logging.Error("QueryAdminPage Error:", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

func GetAdmin(c *gin.Context) {
	id := c.Param("id")

	admin, err := server.GetAdmin(id)

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

func EditAdmin(c *gin.Context) {
	id := c.Param("id")

	admin, err := server.GetAdmin(id)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
		return
	}

	utils.BindJson(c, &admin)

	err = server.EditAdmin(admin)
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

func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	admin, err := server.DeleteAdmin(id)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
	}

	c.JSON(200, gin.H{
		"code": 200,
		"data": admin,
	})
}
