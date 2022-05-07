package api

import (
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/consumer-order/server"
	"github.com/gin-gonic/gin"
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
