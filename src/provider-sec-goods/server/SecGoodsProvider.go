package server

import (
	"context"
	"gimSec/api"
	"gimSec/src/provider-sec-goods/model"
)

type SecGoodsProvider struct {
	api.UnimplementedSecGoodsServerServer
}

func (p *SecGoodsProvider) GetSecGoodsById(ctx context.Context, in *api.SecGoodsId) (*api.SecGoods, error) {
	secGoods, err := model.GetSecGoodsById(in.Id)
	if err != nil {
		return nil, err
	}
	return &api.SecGoods{
		Id:           secGoods.Id,
		Name:         secGoods.Name,
		Price:        secGoods.Price.String(),
		Stock:        secGoods.Stock,
		Photo:        secGoods.Photo,
		Content:      secGoods.Content,
		SecKillStart: secGoods.SecKillStart,
		SecKillEnd:   secGoods.SecKillEnd,
	}, nil
}
