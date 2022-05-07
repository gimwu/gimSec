package server

import (
	"gimSec/api"
	"gimSec/src/consumer-order/model"
)

func AddOrder(userId string, goodss []*api.Goods) (*model.Order, error) {
	return model.AddOrder(userId, goodss)
}
