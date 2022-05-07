package server

import (
	"context"
	"gimSec/api"
)

var GoodsConsumer = new(api.GoodsServiceClientImpl)

func GetGoodsByIds(ids []string) ([]*api.Goods, error) {
	req := api.GoodsIds{Id: ids}
	goods, err := GoodsConsumer.GetGoodsByIds(context.Background(), &req)
	if err != nil {
		return nil, err
	}
	return goods.Goods, nil
}
