package server

import (
	"gimSec/api"
	"gimSec/src/consumer-order/model"
)

func AddOrder(userId string, goodss []*api.Goods) (*model.Order, error) {
	return model.AddOrder(userId, goodss)
}

func GetOrder(id string) (*model.Order, error) {
	return model.GetOrder(id)
}

func Payment(order *model.Order) error {
	return model.UpdateOrder(order)
}

func QueryOrderPage(params map[string]string, currentPage int, pageSize int) (map[string]interface{}, error) {
	orderList, err := model.QueryOrderPage(params, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QueryOrderCount(params)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &orderList
	res["count"] = count
	return res, nil
}
