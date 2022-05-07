package server

import (
	"encoding/json"
	"gimSec/basic/utils"
	"gimSec/src/consumer-sec-order/model"
	"github.com/streadway/amqp"
)

func AddSecOrder(order *model.SecOrder) error {
	order.Id = utils.SnowFlake.NextVal()
	err := model.AddSecOrder(order)
	return err
}

func AddSecOrderByMq(delivery amqp.Delivery) {
	var order model.SecOrder
	err := json.Unmarshal(delivery.Body, &order)
	if err != nil {
		return
	}
	order.Id = utils.SnowFlake.NextVal()
	err = model.AddSecOrder(&order)
	return
}

func QuerySecOrderPage(params interface{}, currentPage int, pageSize int) (map[string]interface{}, error) {
	secOrderList, err := model.QuerySecOrderPage(params, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QuerySecOrderCount(params)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &secOrderList
	res["count"] = count
	return res, nil
}
