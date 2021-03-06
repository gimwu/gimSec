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

func QuerySecOrderPage(params map[string]string, currentPage int, pageSize int) ([]*SecOrder, error) {
	var secOrderList []*SecOrder
	tx := global.DB.Model(&SecOrder{})

	if params["userId"] != "" {
		tx.Where("user_id = ?", params["userId"])
	}
	err := tx.Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&secOrderList).Error
	if err != nil {
		return nil, err
	}
	return secOrderList, nil
}

func QuerySecOrderCount(params map[string]string) (int64, error) {
	var count int64
	tx := global.DB.Model(&SecOrder{})

	if params["userId"] != "" {
		tx.Where("user_id = ?", params["userId"])
	}
	err := tx.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
