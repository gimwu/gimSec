package model

import (
	"gimSec/basic/model"
	"github.com/shopspring/decimal"
)

type OrderStatusEnum int

const (
	DELETE        OrderStatusEnum = iota //删除
	ORDER_NOT_PAY                        //下单未付款
	ORDER_PAY                            //下单已付款
	SHIPED                               //发货
	REVEIVE                              //签收
	REFUND                               //退货
	SUCCESS                              //完成订单
	FAIL                                 //订单失败(未付款自动取消，人为取消)
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

	OrderStatus OrderStatusEnum
}
