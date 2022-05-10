package server

import (
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/model"
)

func AddSecGoods(secGoods *model.SecGoods) error {
	secGoods.Id = utils.SnowFlake.NextVal()
	return model.AddSecOrder(secGoods)
}

func GetSecGoods(id string) (*model.SecGoods, error) {
	return model.GetSecGoodsById(id)
}

func EditSecGoods(secGoods *model.SecGoods) error {
	return model.EditSecGoods(secGoods)
}

func DeleteSecGoods(id string) error {
	secGoods, err := model.GetSecGoodsById(id)
	if err != nil {
		logging.Error(err)
	}
	return model.DeleteSecGoods(secGoods)
}

func QuerySecGoodsPage(order string, currentPage int, pageSize int) (map[string]interface{}, error) {
	secGoodsList, err := model.QuerySecGoodsPage(order, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QuerySecGoodsCount(order)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &secGoodsList
	res["count"] = count
	return res, nil
}
