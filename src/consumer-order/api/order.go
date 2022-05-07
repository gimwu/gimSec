package api

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/consumer-order/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddOrder(c *gin.Context) {
	//TODO 获取用户ID
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	userId := claims.UserId
	// 获取商品信息
	params := make(map[string]interface{}, 0)
	utils.BindJson(c, &params)

	goodsIds := params["goodsIds"]
	slice := utils.ToStringSlice(goodsIds)
	goodsList, err := server.GetGoodsByIds(*slice)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	order, err := server.AddOrder(userId, goodsList)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, order, "下单成功")
}

func QueryOrderPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	logging.Debug(currentPage, pageSize)
	data, err := server.QueryOrderPage(nil, currentPage, pageSize)

	if err != nil {
		logging.Error("QueryUserPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data)
}
