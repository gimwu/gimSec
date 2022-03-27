package v1

import (
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
)

//Addorder shoppping to order
func AddOrder(c *gin.Context) {
	var goodsOrderItems *[]string
	c.ShouldBind(&goodsOrderItems)
	get, _ := c.Get("user")
	user := get.(*model.User)
	order, err := server.AddOrder(goodsOrderItems, user.Id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, order)
}
