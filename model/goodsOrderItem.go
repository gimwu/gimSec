package model

import (
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

//GoodsOrderItem 订单项目表
type GoodsOrderItem struct {
	//继承父类
	model.StateFullEntity

	//goodId 商品id
	GoodId string `gorm:"type:varchar(255);not null"`

	//usernameId 购买用户id
	UsernameId string `gorm:"type:varchar(255);not null"`

	//goodsNum 购买商品数量
	GoodsNum uint

	//price 总共价格
	Price decimal.Decimal
}
