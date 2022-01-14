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
	name string

	//price 商品价格
	price decimal.Decimal

	//stock 商品库存
	stock uint64

	//photo 商品图片
	photo string

	//商品描述
	content string

	//商品所属商家id
	belongUsernameId string
}
