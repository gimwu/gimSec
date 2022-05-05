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

func AddSecOrder(goods *SecGoods) error {
	err := global.DB.Create(goods).Error
	if err != nil {
		return err
	}
	return nil
}

func GetSecGoodsById(id string) (*SecGoods, error) {
	var secGoods SecGoods
	err := global.DB.Where("id = ?", id).First(&secGoods).Error
	if err != nil {
		return nil, err
	}
	return &secGoods, nil
}
