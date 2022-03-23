package server

import (
	"gimSec/model"
)

func AddGoodsOrderItem(goods *model.Goods, user *model.User) (bool, error) {
	isExist, err := model.CheckGoodsStock(goods)
	if err != nil {
		return false, err
	}
	if !isExist {
		return false, nil
	}
	goods.Stock--

}
