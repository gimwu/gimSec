package model

import (
	"gimSec/basic/model"
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

func AddOrder(order *Order) error {
	err := db.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}
