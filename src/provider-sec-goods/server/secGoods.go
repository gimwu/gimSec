package server

import (
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/model"
)

func AddSecGoods(secGoods *model.SecGoods) error {
	secGoods.Id = utils.SnowFlake.NextVal()
	return model.AddSecOrder(secGoods)
}
