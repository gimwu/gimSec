package server

import (
	"gimSec/basic/utils"
	"gimSec/model"
)

func AddOrder(goodsOrderItem *[]string, user string) (*model.Order, error) {
	order := &model.Order{}
	order.Id = utils.SnowFlake.NextVal()
	sumPrice, err := model.GetGoodsOrderItemPriceByIds(goodsOrderItem)
	if err != nil {
		return nil, err
	}
	order.Price = *sumPrice
	order.BelongUserId = user
	err = model.AddOrder(order)
	if err != nil {
		return nil, err
	}
	err = model.UpdatesGoodsOrderItem(goodsOrderItem, order.Id)
	if err != nil {
		return nil, err
	}
	return order, nil

}
