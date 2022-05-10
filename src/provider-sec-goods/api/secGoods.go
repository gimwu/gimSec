package api

import (
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/model"
	"gimSec/src/provider-sec-goods/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddSecGoods(c *gin.Context) {
	var secGoods model.SecGoods
	utils.BindJson(c, &secGoods)
	err := server.AddSecGoods(&secGoods)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, secGoods)
}

func EditSecGoods(c *gin.Context) {
	id := c.Query("id")
	logging.Debug(id)

	goods, err := server.GetSecGoods(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	utils.BindJson(c, &goods)

	err = server.EditSecGoods(goods)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, goods, nil)
}

func DeleteSecGoods(c *gin.Context) {
	var params map[string]string
	err := utils.BindJson(c, &params)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	err = server.DeleteSecGoods(params["id"])
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, nil, "success Delete")
}

func QuerySecGoodsPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	order := c.Query("order")
	logging.Debug(currentPage, pageSize)
	data, err := server.QuerySecGoodsPage(order, currentPage, pageSize)
	if err != nil {
		logging.Error("QueryGoodsPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data, nil)
}
