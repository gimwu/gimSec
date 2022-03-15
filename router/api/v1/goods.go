package v1

import (
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/model"
	"gimSec/server"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddGoods(c *gin.Context) {
	json := model.Goods{}
	utils.BindJson(c, &json)
	err := server.AddGoods(&json)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": "500",
			"data": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": "200",
		"data": &json,
		"msg":  "AddGoods Success",
	})
}

func DeleteGoods(c *gin.Context) {
	id := c.Param("id")
	err := server.DeleteGoods(id)
	if err != nil {
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})

	}
}

func EditGoods(c *gin.Context) {
	id := c.Param("id")
	logging.Debug(id)

	goods, err := server.GetGoods(id)
	if err != nil {
		logging.Error(err)
		c.JSON(200, gin.H{
			"msg":  "500",
			"data": err,
		})
		return
	}

	utils.BindJson(c, &goods)

	err = server.EditGoods(goods)
	if err != nil {
		logging.Error(err)
		c.JSON(500, gin.H{
			"code": "500",
			"data": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": "200",
		"data": goods,
		"msg":  "success",
	})
}

//GetGoods select by id
func GetGoods(c *gin.Context) {
	id := c.Param("id")
	logging.Debug(id)

	goods, err := server.GetGoods(id)
	if err != nil {
		logging.Error(err)
		c.JSON(200, gin.H{
			"msg":  "500",
			"data": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":  "200",
		"data": goods,
	})

}

func QueryGoodsPage(c *gin.Context) {
	json := make(map[string]interface{})
	utils.BindJson(c, &json)
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	logging.Debug(currentPage, pageSize)
	data, err := server.QueryGoodsPage(&json, currentPage, pageSize)
	if err != nil {
		logging.Error("QueryGoodsPage Error:", err)
		c.JSON(500, gin.H{
			"code": 500,
			"data": err,
		})
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}
