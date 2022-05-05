package api

import (
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/model"
	"gimSec/src/provider-sec-goods/server"
	"github.com/gin-gonic/gin"
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
