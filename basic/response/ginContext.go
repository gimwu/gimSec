package response

import (
	"github.com/gin-gonic/gin"
)

type HTTPStatusNum int

const (
	SUCCESS          HTTPStatusNum = 200
	ERROR            HTTPStatusNum = 500
	PERMISSION_ERROR HTTPStatusNum = 403
	STOCK_ERROR      HTTPStatusNum = 5677
)

func Success(c *gin.Context, data interface{}, msg ...interface{}) {
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"data": data,
		"msg":  msg,
	})
}

func Error(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": ERROR,
		"data": data,
	})
	c.Abort()
}

func Info(c *gin.Context, code HTTPStatusNum, data interface{}) {
	c.JSON(int(code), gin.H{
		"code": code,
		"data": data,
	})
}
