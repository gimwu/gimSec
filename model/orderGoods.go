package model

import (
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//OrderGoods 订单项目关联表
type OrderGoods struct {
	//继承父类
	model.StateFullEntity

	//orderId 订单id
	orderId string

	//orderItemId 订单项目id
	orderItemId string

	//orderItemPrice 项目总价
	orderItemPrice decimal.Decimal
}
