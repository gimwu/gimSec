package model

import (
	"gimSec/api"
	"gimSec/basic/global"
	"gimSec/basic/model"
	"gimSec/basic/utils"
	"github.com/shopspring/decimal"
)

type OrderStatusEnum int

const (
	_       = iota
	NOT_PAY //未付款
	PAY     //付款未发货
	SEND    //发货
	SUCCESS //成功到货
	FAIL    //取消订单
)

//Order 订单表
type Order struct {
	//继承父类
	model.StateFullEntity

	//price 订单总价
	Price decimal.Decimal

	//OrderStatus 订单状态 1待付款 2待发货 3待收获 4已完成
	OrderStatus OrderStatusEnum `gorm:"type:int;not null"`

	//所属用户id
	BelongUserId string `gorm:"type:varchar(255);not null"`

	Addr string `gorm:"type:varchar(255);not null""`
}

func AddOrder(userId string, goodss []*api.Goods) (*Order, error) {
	tx := global.DB.Begin()
	defer tx.Rollback()
	orderId := utils.SnowFlake.NextVal()

	price := decimal.Decimal{}
	orderItems := make([]OrderGoods, 0)
	goodsIds := make([]string, 0)
	for _, goods := range goodss {
		fromString, _ := decimal.NewFromString(goods.Price)
		price = decimal.Sum(price, fromString)

		orderGoods := &OrderGoods{
			OrderId:        orderId,
			OrderItemId:    goods.Id,
			OrderItemPrice: decimal.Decimal{},
		}
		goodsIds = append(goodsIds, goods.Id)
		orderGoods.Id = utils.SnowFlake.NextVal()
		orderItems = append(orderItems, *orderGoods)

	}
	order := &Order{
		Price:        price,
		OrderStatus:  1,
		BelongUserId: userId,
	}
	order.Id = orderId

	err := tx.Exec("update goods set stock = stock -1 where id in ?", goodsIds).Error
	if err != nil {
		return nil, err
	}

	err = tx.Model(&Order{}).Create(&order).Error
	if err != nil {
		return nil, err
	}

	err = tx.Model(&OrderGoods{}).Create(&orderItems).Error
	if err != nil {
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func QueryOrderPage(params map[string]string, currentPage int, pageSize int) ([]*Order, error) {
	var orderList []*Order
	tx := global.DB.Model(&Order{})

	if params["belongUserId"] != "" {
		tx.Where("belong_user_id = ?", params["belongUserId"])
	}
	err := tx.Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&orderList).Error
	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func QueryOrderCount(params map[string]string) (int64, error) {
	var count int64
	tx := global.DB.Model(&Order{})

	if params["belongUserId"] != "" {
		tx.Where("belong_user_id = ?", params["belongUserId"])
	}
	err := tx.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetOrder(id string) (*Order, error) {
	var order Order
	err := global.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func UpdateOrder(order *Order) error {
	if err := global.DB.Updates(order).Error; err != nil {
		return err
	}
	return nil
}
