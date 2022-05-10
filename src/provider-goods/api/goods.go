package v1

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/provider-goods/model"
	"gimSec/src/provider-goods/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddGoods(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	userId := claims.UserId

	json := model.Goods{}
	err = utils.BindJson(c, &json)
	json.BelongUsernameId = userId
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	err = server.AddGoods(&json)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, json, "success Delete")
}

func DeleteGoods(c *gin.Context) {
	var params map[string]string
	err := utils.BindJson(c, &params)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	err = server.DeleteGoods(params["id"])
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, nil, "success Delete")
}

func EditGoods(c *gin.Context) {
	id := c.Query("id")
	logging.Debug(id)

	goods, err := server.GetGoods(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	utils.BindJson(c, &goods)

	err = server.EditGoods(goods)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, goods, nil)
}

//GetGoods select by id
func GetGoods(c *gin.Context) {
	id := c.Query("id")
	logging.Debug(id)

	goods, err := server.GetGoods(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, goods, nil)

}

func QueryGoodsPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	order := c.Query("order")
	logging.Debug(currentPage, pageSize)
	data, err := server.QueryGoodsPage(order, currentPage, pageSize)
	if err != nil {
		logging.Error("QueryGoodsPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data, nil)
}
