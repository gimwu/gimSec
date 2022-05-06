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
	var secGoods *api.SecGoods
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
	}
	if err != nil && err != redis.Nil {
		logging.Error(err.Error())
		response.Error(c, err)
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