package server

import (
	"gimSec/basic/utils"
	"gimSec/src/provider-sec-goods/model"
)

func AddSecGoods(secGoods *model.SecGoods) error {
	secGoods.Id = utils.SnowFlake.NextVal()
	return model.AddSecOrder(secGoods)
}

func QuerySecGoodsPage(params interface{}, currentPage int, pageSize int) (map[string]interface{}, error) {
	secGoodsList, err := model.QuerySecGoodsPage(params, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QuerySecGoodsCount(params)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &secGoodsList
	res["count"] = count
	return res, nil
}
