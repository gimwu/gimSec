package server

import (
	"gimSec/basic/utils"
	"gimSec/src/consumer-sec-order/model"
)

func AddSecOrder(order *model.SecOrder) error {
	order.Id = utils.SnowFlake.NextVal()
	err := model.AddSecOrder(order)
	return err
}
