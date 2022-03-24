package utils

import (
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"github.com/gin-gonic/gin"
)

func BindJson(c *gin.Context, obj interface{}) {
	err := c.BindJSON(obj)
	if err != nil {
		logging.Error("BindJson Error:", err)
		response.Error(c, err.Error())
		return
	}
}
