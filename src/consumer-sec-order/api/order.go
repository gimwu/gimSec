package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"gimSec/basic/global"
	"gimSec/basic/jwt"
	"gimSec/basic/logging"
	"gimSec/basic/response"
	"gimSec/basic/utils"
	"gimSec/src/consumer-sec-order/model"
	"gimSec/src/consumer-sec-order/server"
	"github.com/gin-gonic/gin"
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
		response.Info(c, 202, "已经参与秒杀")
		return
	}
	//TODO 商品是否有库存

	secGoods, err := model.GetSecGoodsById(goodsId)
	if err != nil {
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

	err = server.AddSecOrder(secGoods)
	if err != nil {
		response.Error(c, err.Error())
		logging.Error(err)
		return
	}

	order := &model.SecOrder{
		UserId:      userId,
		GoodsId:     goodsId,
		Price:       secGoods.Price,
		OrderStatus: 1,
	}

	//TODO 异步处理，生成订单
	marshal, err := json.Marshal(order)
	err = global.CH.PublishToQueue(marshal, "order")
	if err != nil {
		response.Error(c, err.Error())
		logging.Error(err)
		return
	}
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
