package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"gimSec/api"
	"gimSec/basic/global"
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/consumer-sec-order/model"
	"gimSec/src/consumer-sec-order/server"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/shopspring/decimal"
	"strconv"
	"time"
)

func AddSecOrder(c *gin.Context) {
	//TODO 获取用户ID
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	userId := claims.UserId

	//TODO 获取商品信息
	params := make(map[string]interface{}, 0)
	utils.BindJson(c, &params)

	goodsId := params["goodsId"].(string)
	//TODO 用户是否正在秒杀
	goodsIDAndUserId := fmt.Sprintf("%s%s", goodsId, userId)
	b, err := global.REDIS.Exists(context.Background(), goodsIDAndUserId).Result()
	if err != nil {
		logging.Error(err.Error())
		response.Error(c, err)
		return
	}
	if b == 1 {
		response.Info(c, 200, "已经参与秒杀")
		return
	}
	//TODO 商品是否有库存

	//TODO 判断商品是否已经存入缓存中 如果不存在则写入缓存中
	goodsStock1, err := global.REDIS.HGet(context.Background(), goodsId, "stock").Result()
	secGoods := &api.SecGoods{}
	secGoodsMap := make(map[string]string, 0)
	if err == redis.Nil {
		var err2 error
		secGoods, err2 = server.GetGoodsById(goodsId)
		if err2 != nil {
			logging.Error(err2)
			response.Error(c, err2.Error())
			return
		}
		global.REDIS.HSet(context.Background(), goodsId,
			"id", secGoods.Id,
			"name", secGoods.Name,
			"price", secGoods.Price,
			"stock", secGoods.Stock,
			"photo", secGoods.Photo,
			"content", secGoods.Content,
			"secKillStart", secGoods.SecKillStart,
			"secKillEnd", secGoods.SecKillEnd)

		secGoodsMap["id"] = secGoods.Id
		secGoodsMap["name"] = secGoods.Name
		secGoodsMap["price"] = secGoods.Price
		secGoodsMap["photo"] = secGoods.Photo
		secGoodsMap["content"] = secGoods.Content
		secGoodsMap["stock"] = strconv.FormatInt(secGoods.Stock, 10)
		secGoodsMap["secKillStart"] = strconv.FormatInt(secGoods.SecKillStart, 10)
		secGoodsMap["secKillEnd"] = strconv.FormatInt(secGoods.SecKillEnd, 10)
		goodsStock1 = strconv.FormatInt(secGoods.Stock, 10)
	} else {
		var err2 error
		secGoodsMap, err2 = global.REDIS.HGetAll(context.Background(), goodsId).Result()
		if err2 != nil {
			logging.Error(err.Error())
			response.Error(c, err)
			return
		}
		secGoods.Id = secGoodsMap["id"]
		secGoods.Name = secGoodsMap["name"]
		secGoods.Price = secGoodsMap["price"]
		stock, _ := strconv.ParseInt(secGoodsMap["stock"], 10, 64)
		secGoods.Stock = stock
		secGoods.Photo = secGoodsMap["photo"]
		secGoods.Content = secGoodsMap["content"]
		secKillStart, _ := strconv.ParseInt(secGoodsMap["secKillStart"], 10, 64)
		secGoods.SecKillStart = secKillStart
		secKillEnd, _ := strconv.ParseInt(secGoodsMap["secKillEnd"], 10, 64)
		secGoods.SecKillEnd = secKillEnd
	}
	if err != nil && err != redis.Nil {
		logging.Error(err.Error())
		response.Error(c, err)
		return
	}
	//TODO 未到达时间 直接返回
	now := time.Now().Unix()
	if now < secGoods.SecKillStart || now > secGoods.SecKillEnd {
		response.Info(c, 201, "未达到开始活动时间")
		logging.Info("未达到开始活动时间")
		return
	}

	goodsStock, err := strconv.Atoi(goodsStock1)

	//库存小于0 则直接返回
	if goodsStock <= 0 {
		response.Info(c, 200, "商品已抢购一空")
		return
	}
	//TODO 预扣库存
	result, err := global.REDIS.SetNX(context.Background(), goodsIDAndUserId, "1", 5*time.Second).Result()
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	if result != true {
		logging.Error("预扣库存加锁失败")
		response.Error(c, "预扣库存加锁失败")
		return
	}
	all, err := global.REDIS.HGetAll(context.Background(), goodsId).Result()

	order := &model.SecOrder{
		UserId:      userId,
		GoodsId:     goodsId,
		Price:       decimal.RequireFromString(all["price"]),
		OrderStatus: 1,
	}
	global.REDIS.HIncrBy(context.Background(), goodsId, "stock", -1)
	//TODO 异步处理，生成订单
	marshal, err := json.Marshal(order)
	err = global.CH.PublishToQueue(marshal, "order")
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	//TODO 结束

	global.REDIS.Del(context.Background(), goodsIDAndUserId)
	response.Success(c, "200", "秒杀成功")
}

func QuerySecOrderPage(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	logging.Debug(currentPage, pageSize)
	data, err := server.QuerySecOrderPage(nil, currentPage, pageSize)

	if err != nil {
		logging.Error("QueryUserPage Error:", err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data)
}

func QueryMyOrder(c *gin.Context) {
	currentPage, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	params := c.QueryMap("params")
	tokenString := c.GetHeader("Authorization")
	tokenString = tokenString[7:]
	_, claims, err := jwt.ParseToken(tokenString)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	userId := claims.UserId
	params["userId"] = userId

	data, err := server.QuerySecOrderPage(params, currentPage, pageSize)
	if err != nil {
		logging.Error(err)
		response.Error(c, err.Error())
		return
	}
	response.Success(c, data)
}
