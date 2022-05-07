package model

import (
	"gimSec/api"
	"gimSec/basic/global"
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//OrderGoods 订单项目关联表
type OrderGoods struct {
	//继承父类
	model.StateFullEntity

	//orderId 订单id
	OrderId string `gorm:"type:varchar(255);not null"`

	//orderItemId 订单项目id
	OrderItemId string `gorm:"type:varchar(255);not null"`

	//orderItemPrice 项目总价
	OrderItemPrice decimal.Decimal
}

func AddOrderGoods(orderId string, goodsList []api.Goods) error {
	var orderGoodsList []OrderGoods
	for _, s := range goodsList {
		fromString, _ := decimal.NewFromString(s.Price)
		orderGoods := &OrderGoods{
			OrderId:        orderId,
			OrderItemId:    s.Id,
			OrderItemPrice: fromString,
		}
		orderGoodsList = append(orderGoodsList, *orderGoods)
	}
	err := global.DB.Create(&orderGoodsList).Error
	return err
}
