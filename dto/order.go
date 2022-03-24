package dto

import "gimSec/model"

type order struct {
	orderId         string                  `json:"OrderId"`
	goodsOrderItems []*model.GoodsOrderItem `json:"goodsOrderItems"`
}
