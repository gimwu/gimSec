package model

import (
	"gimSec/api"
	"gimSec/basic/global"
	"gimSec/basic/model"
	"gimSec/basic/utils"
	"github.com/shopspring/decimal"
)

//Order 订单表
type Order struct {
	//继承父类
	model.StateFullEntity

	//price 订单总价
	Price decimal.Decimal

	//OrderStatus 订单状态 1待付款 2待发货 3待收获 4已完成
	OrderStatus int `gorm:"type:int;not null"`

	//所属用户id
	BelongUserId string `gorm:"type:varchar(255);not null"`
}

func AddOrder(userId string, goodss []*api.Goods) (*Order, error) {
	tx := global.DB.Begin()
	defer tx.Rollback()
	orderId := utils.SnowFlake.NextVal()

	price := decimal.Decimal{}
	orderItems := make([]OrderGoods, 0)
	for _, goods := range goodss {
		fromString, _ := decimal.NewFromString(goods.Price)
		price = decimal.Sum(price, fromString)

		orderGoods := &OrderGoods{
			OrderId:        orderId,
			OrderItemId:    goods.Id,
			OrderItemPrice: decimal.Decimal{},
		}
		orderGoods.Id = utils.SnowFlake.NextVal()
		orderItems = append(orderItems, *orderGoods)

	}
	order := &Order{
		Price:        price,
		OrderStatus:  1,
		BelongUserId: userId,
	}
	order.Id = orderId

	err := tx.Model(&Order{}).Create(&order).Error
	if err != nil {
		return nil, err
	}

	err = tx.Model(&OrderGoods{}).Create(&orderItems).Error
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return order, nil
}
