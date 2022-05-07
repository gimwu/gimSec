package model

import (
	"gimSec/basic/global"
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//SecOrder 订单表
type SecOrder struct {
	//继承父类
	model.StateFullEntity

	UserId string `gorm:"varchar(255);not null"`

	GoodsId string `gorm:"varchar(255);not null"`

	//price 订单总价
	Price decimal.Decimal

	//OrderStatus 订单状态 1待付款 2待发货 3待收获 4已完成
	OrderStatus int `gorm:"type:int;not null"`
}

func AddSecOrder(order *SecOrder) error {
	err := global.DB.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func QuerySecOrderPage(params interface{}, currentPage int, pageSize int) ([]*SecOrder, error) {
	var secOrderList []*SecOrder
	err := global.DB.Model(&SecOrder{}).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&secOrderList).Error
	if err != nil {
		return nil, err
	}
	return secOrderList, nil
}

func QuerySecOrderCount(params interface{}) (int64, error) {
	var count int64
	err := global.DB.Model(&SecOrder{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
