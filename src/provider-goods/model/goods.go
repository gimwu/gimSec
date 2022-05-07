package model

import (
	"gimSec/basic/global"
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//Goods 普通商品表
type Goods struct {
	// 继承父类
	model.StateFullEntity

	//name 商品名称
	Name string `gorm:"type:varchar(255);not null"`

	//price 商品价格
	Price decimal.Decimal

	//stock 商品库存
	Stock int64

	//photo 商品图片
	Photo string `gorm:"type:varchar(255);not null"`

	//商品描述
	Content string `gorm:"type:varchar(255);not null"`

	//商品所属商家id
	BelongUsernameId string `gorm:"type:varchar(255);not null"`
}

func GetGoods(id string) (*Goods, error) {
	var goods Goods
	err := global.DB.Where("id = ?", id).First(&goods).Error
	if err != nil {
		return nil, err
	}
	return &goods, nil
}

func GetGoodsByIds(ids []string) ([]Goods, error) {
	goodsList := make([]Goods, 0)
	err := global.DB.Where("id in ?", ids).Find(&goodsList).Error
	if err != nil {
		return nil, err
	}
	return goodsList, nil
}

func AddGoods(goods *Goods) error {
	if err := global.DB.Create(&goods).Error; err != nil {
		return err
	}
	return nil
}

func EditGoods(goods *Goods) error {
	if err := global.DB.Updates(goods).Error; err != nil {
		return err
	}
	return nil
}

func DeleteGoods(goods *Goods) error {
	if err := global.DB.Delete(&goods).Error; err != nil {
		return err
	}
	return nil
}

func QueryGoodsPage(params interface{}, currentPage int, pageSize int) ([]*Goods, error) {
	var GoodsList []*Goods
	err := global.DB.Model(&Goods{}).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&GoodsList).Error
	if err != nil {
		return nil, err
	}
	return GoodsList, nil
}

func QueryGoodsCount(params interface{}) (int64, error) {
	var count int64
	err := global.DB.Model(&Goods{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CheckGoodsStock(goods *Goods) (bool, error) {
	var stock int
	err := global.DB.Model(goods).Select("stock").Where("id = ?", goods.Id).First(&stock).Error
	if err != nil {
		return false, err
	}
	if stock > 0 {
		return true, nil
	}
	return false, nil
}
