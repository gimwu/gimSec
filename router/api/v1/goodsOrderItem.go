package v1

import (
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
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
