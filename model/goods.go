package model

import (
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//Goods 普通商品表
type Goods struct {
	// 继承父类
	model.StateFullEntity

	//name 商品名称
	Name string

	//price 商品价格
	Price decimal.Decimal

	//stock 商品库存
	Stock uint64

	//photo 商品图片
	Photo string

	//商品描述
	Content string

	//商品所属商家id
	BelongUsernameId string
}
