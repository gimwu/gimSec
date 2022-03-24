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

	//orderId 订单id
	OrderId string `gorm:"type:varchar(255);not null"`

	//usernameId 购买用户id
	UsernameId string `gorm:"type:varchar(255);not null"`

	//goodsNum 购买商品数量
	GoodsNum int64

	//price 总共价格
	Price decimal.Decimal

	OrderStatus OrderStatusEnum
}

func AddGoodsOrderItem(goodsOrderItem *GoodsOrderItem) error {
	err := db.Create(goodsOrderItem).Error
	if err != nil {
		return err
	}
	return nil
}

func GetGoodsOrderItem(id string) (*GoodsOrderItem, error) {
	var goodsOrderitem GoodsOrderItem
	err := db.Where("id = ?", id).First(&goodsOrderitem).Error
	if err != nil {
		return nil, err
	}
	return &goodsOrderitem, nil
}

func DeleteGoodsOrderItem(goodsOrderItem *GoodsOrderItem) (*GoodsOrderItem, error) {
	err := db.Delete(&goodsOrderItem).Error
	if err != nil {
		return nil, err
	}
	return goodsOrderItem, nil
}

func QueryGoodsOrderItemPage(params interface{}, currentPage int, pageSize int) ([]*GoodsOrderItem, error) {
	var goodsOrderItemList []*GoodsOrderItem
	err := db.Model(&GoodsOrderItem{}).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&goodsOrderItemList).Error
	if err != nil {
		return nil, err
	}
	return goodsOrderItemList, nil
}

func QueryGoodsOrderItemCount(params interface{}) (int64, error) {
	var count int64
	err := db.Model(&GoodsOrderItem{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil

}

func GetGoodsOrderItemPriceByIds(goodsOrderItems *[]string) (*decimal.Decimal, error) {
	var price decimal.Decimal
	err := db.Select("sum(price)").Model(&GoodsOrderItem{}).First(&price).Where("id in ?", goodsOrderItems).Error
	if err != nil {
		return &decimal.Zero, err
	}
	return &price, nil
}

func UpdatesGoodsOrderItem(goodsOrderItems *[]string, orderId string) error {
	err := db.Model(&GoodsOrderItem{}).Where("id in ?", goodsOrderItems).Updates(GoodsOrderItem{OrderId: orderId}).Error
	if err != nil {
		return err
	}
	return nil
}
