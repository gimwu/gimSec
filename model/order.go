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

	//所属用户id
	BelongUserId string `gorm:"type:varchar(255);not null"`
}
