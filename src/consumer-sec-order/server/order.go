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
