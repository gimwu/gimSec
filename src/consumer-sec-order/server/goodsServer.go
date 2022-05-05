package server

import (
	"context"
	"gimSec/api"
)

var SecGoodsConsumer = new(api.SecGoodsServerClientImpl)

func GetGoodsById(id string) (*api.SecGoods, error) {
	req := &api.SecGoodsId{Id: id}
	secGoods, err := SecGoodsConsumer.GetSecGoodsById(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return secGoods, nil
}
