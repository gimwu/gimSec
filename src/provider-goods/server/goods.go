package server

import (
	"gimSec/basic/logging"
	"gimSec/basic/utils"
	"gimSec/src/provider-goods/model"
)

func GetGoods(id string) (*model.Goods, error) {
	return model.GetGoods(id)
}

func SelectGoodsByIds(ids []string) ([]model.Goods, error) {
	return model.GetGoodsByIds(ids)
}

func AddGoods(goods *model.Goods) error {
	goods.Id = utils.SnowFlake.NextVal()
	return model.AddGoods(goods)
}

func EditGoods(goods *model.Goods) error {
	return model.EditGoods(goods)
}

func DeleteGoods(id string) error {
	goods, err := model.GetGoods(id)
	if err != nil {
		logging.Error(err)
	}
	return model.DeleteGoods(goods)
}

func QueryGoodsPage(params interface{}, currentPage int, pageSize int) (map[string]interface{}, error) {
	GoodsList, err := model.QueryGoodsPage(params, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QueryGoodsCount(params)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &GoodsList
	res["count"] = count
	return res, nil
}
