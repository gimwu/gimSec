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

func QuerySecGoodsPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	logging.Debug(currentPage, pageSize)
	data, err := server.QuerySecGoodsPage(nil, currentPage, pageSize)
	if err != nil {
		logging.Error("QueryGoodsPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data, nil)
}
