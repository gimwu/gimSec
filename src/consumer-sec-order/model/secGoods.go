package model

import (
	"gimSec/basic/global"
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

type SecGoods struct {
	model.StateFullEntity

	Name string `gorm:"type:varchar(255);not null"`

	Price decimal.Decimal `gorm:""`

	Photo string `gorm:"type:varchar(255);not null"`

	Stock int64 `gorm:"type:int;not null"`

	Content string `gorm:"type:varchar(255);not null"`

	SecKillStart int64 `gorm:"type:int;not null" json:"secKillStart"`

	SecKillEnd int64 `gorm:"type:int;not null" json:"secKillEnd"`
}

func UpdateSecGoods(goods *SecGoods) error {
	err := global.DB.Exec("update sec_goods set stock = stock-1 where id = ?", goods.Id).Error
	if err != nil {
		return err
	}
	return nil
}

func AddSecOrder(order *SecOrder) error {
	err := global.DB.Create(order).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteSecGoods(secGoods *SecGoods) error {
	if err := global.DB.Delete(&secGoods).Error; err != nil {
		return err
	}
	return nil
}

func QuerySecGoodsPage(order string, currentPage int, pageSize int) ([]*SecGoods, error) {
	var secGoodsList []*SecGoods
	global.DB.Model(&SecGoods{}).Offset((currentPage - 1) * pageSize).Limit(pageSize)
	if order != "" {
		global.DB.Order(order)
	} else {
		global.DB.Order("create_at")
	}
	err := global.DB.Find(&secGoodsList).Error
	if err != nil {
		return nil, err
	}
	return secGoodsList, nil
}

func QuerySecGoodsCount(params interface{}) (int64, error) {
	var count int64
	err := global.DB.Model(&SecGoods{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetSecGoodsById(id string) (*SecGoods, error) {
	var secGoods SecGoods
	err := global.DB.Where("id = ?", id).First(&secGoods).Error
	if err != nil {
		return nil, err
	}
	return &secGoods, nil
}

func EditSecGoods(secGoods *SecGoods) error {
	if err := global.DB.Updates(secGoods).Error; err != nil {
		return err
	}
	return nil
}
