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
	goods, err := server.GetGoods(goodsId)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	logging.Info(goods)
	get, _ := c.Get("user")
	user := get.(model.User)
	server.AddGoodsOrderItem(goods, &user)
	logging.Debug(user.Id, "       ", user.Name)
}
