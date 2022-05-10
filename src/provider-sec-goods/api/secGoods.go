package api

import (
	"context"
	"encoding/json"
	"fmt"
	"gimSec/basic/global"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/model"
	"gimSec/src/provider-sec-goods/server"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
)

func GetSecGoods(c *gin.Context) {
	id := c.Query("id")

	result, err := global.REDIS.Get(context.Background(), id).Result()
	if err == redis.Nil {
		goods, err := server.GetSecGoods(id)
		if err != nil {
			logging.Error(err)
			response.Error(c, err.Error())
			return
		}
		marshal, _ := json.Marshal(goods)
		_, err = global.REDIS.Set(context.Background(), id, marshal, redis.KeepTTL).Result()
		if err != nil {
			response.Error(c, err.Error())
			logging.Error(err)
			return
		}
		response.Success(c, goods, nil)
		return
	}
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	var res map[string]interface{}
	err2 := json.Unmarshal([]byte(result), &res)
	if err2 != nil {
		logging.Error(err2)
		response.Error(c, err2.Error())
		return
	}

	response.Success(c, res)
}

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

func EditSecGoods(c *gin.Context) {
	id := c.Query("id")
	logging.Debug(id)

	goods, err := server.GetSecGoods(id)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	utils.BindJson(c, &goods)

	err = server.EditSecGoods(goods)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}

	response.Success(c, goods, nil)
}

func DeleteSecGoods(c *gin.Context) {
	var params map[string]string
	err := utils.BindJson(c, &params)
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	err = server.DeleteSecGoods(params["id"])
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, nil, "success Delete")
}

func QuerySecGoodsPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	order := c.Query("order")
	pageName := fmt.Sprintf("%s%d%d", "secGoodsPage", currentPage, pageSize)
	logging.Debug(currentPage, pageSize)
	result, err2 := global.REDIS.Get(context.Background(), pageName).Result()
	if err2 == redis.Nil {
		data, err := server.QuerySecGoodsPage(order, currentPage, pageSize)
		if err != nil {
			logging.Error("QueryGoodsPage Error:", err)
			response.Error(c, err.Error())
			return
		}
		marshal, _ := json.Marshal(data)
		_, err = global.REDIS.Set(context.Background(), pageName, marshal, redis.KeepTTL).Result()
		if err != nil {
			response.Error(c, err.Error())
			logging.Error(err)
			return
		}
		response.Success(c, data, nil)
		return
	}
	var res map[string]interface{}
	err2 = json.Unmarshal([]byte(result), &res)
	if err2 != nil {
		logging.Error(err2)
		response.Error(c, err2.Error())
		return
	}

	response.Success(c, res)

}
