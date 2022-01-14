package model

import (
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//GoodsOrderItem 订单项目表
type GoodsOrderItem struct {
	//继承父类
	model.StateFullEntity

	//goodId 商品id
	goodId string

	//usernameId 购买用户id
	usernameId string

	//goodsNum 购买商品数量
	goodsNum uint

	//price 总共价格
	price decimal.Decimal
}