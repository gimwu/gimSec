package utils

import (
	"gimSec/basic/logging"
	"github.com/gin-gonic/gin"
)

func BindJson(c *gin.Context, obj interface{}) {
	err := c.BindJSON(obj)
	if err != nil {
		logging.Error("BindJson Error:", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
			"msg":  "BindJson Error",
		})

	}
}
