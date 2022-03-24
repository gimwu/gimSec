package v1

import (
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddGoodsOrderItem(c *gin.Context) {
	json := make(map[string]interface{})
	utils.BindJson(c, &json)

	goodsId := json["goodsId"].(string)
	goodsNum := int64(json["goodsNum"].(float64))
	goods, err := server.GetGoods(goodsId)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	logging.Info(goods)
	get, _ := c.Get("user")
	user := get.(*model.User)
	goodsOrderItem, err := server.AddGoodsOrderItem(goods, goodsNum, user)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, goodsOrderItem)
}

func DeleteGoodsOrderItem(c *gin.Context) {
	id := c.Query("id")
	goodsOrderItem, err := server.DeleteGoodsOrderItem(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, goodsOrderItem)
}

func GetGoodsOrderItem(c *gin.Context) {
	id := c.Query("id")
	goodsOrderItem, err := server.GetGoodsOrderItem(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, goodsOrderItem)
}

func QueryGoodsOrderItem(c *gin.Context) {
	json := make(map[string]interface{})
	utils.BindJson(c, &json)
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	data, err := server.QueryGoodsOrderItemPage(&json, currentPage, pageSize)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data, nil)
}
