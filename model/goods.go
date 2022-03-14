package model

import (
	"gimSec/basic/model"
	"gimSec/basic/utils"
	"github.com/shopspring/decimal"
)

var db = utils.GormMysqlDatabase

//Goods 普通商品表
type Goods struct {
	// 继承父类
	model.StateFullEntity

	//name 商品名称
	Name string `gorm:"type:varchar(255);not null"`

	//price 商品价格
	Price decimal.Decimal

	//stock 商品库存
	Stock uint64

	//photo 商品图片
	Photo string `gorm:"type:varchar(255);not null"`

	//商品描述
	Content string `gorm:"type:varchar(255);not null"`

	//商品所属商家id
	BelongUsernameId string `gorm:"type:varchar(255);not null"`
}

func GetGoods(id string) (*Goods, error) {
	var goods Goods
	err := db.Where("id = ?", id).First(&goods).Error
	if err != nil {
		return nil, err
	}
	return &goods, nil
}

func AddGoods(goods *Goods) error {
	if err := db.Create(&goods).Error; err != nil {
		return err
	}
	return nil
}

func EditGoods(goods *Goods) error {
	if err := db.Updates(goods).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGoods(goods *Goods) error {
	if err := db.Delete(&goods).Error; err != nil {
		return err
	}
	return nil
}

func QueryGoodsPage(params interface{}, currentPage int, pageSize int) ([]*Goods, error) {
	var GoodsList []*Goods
	err := db.Where(params).Offset(currentPage).Limit(pageSize).Find(&GoodsList).Error
	if err != nil {
		return nil, err
	}
	return GoodsList, nil
}

func QueryGoodsCount(params interface{}) (int64, error) {
	var count int64
	err := db.Model(&Goods{}).Where(params).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
