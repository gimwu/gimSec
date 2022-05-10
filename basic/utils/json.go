package utils

import (
	"github.com/gin-gonic/gin"
)

func BindJson(c *gin.Context, obj interface{}) error {
	err := c.BindJSON(obj)
	if err != nil {
		return err
	}
	return nil
}
