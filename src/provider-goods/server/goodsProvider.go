package server

import (
	"context"
	"gimSec/api"
	"gimSec/src/provider-goods/model"
)

type GoodsProvider struct {
	api.UnimplementedGoodsServiceServer
}

func (p *GoodsProvider) GetGoodsById(ctx context.Context, in *api.GoodsId) (*api.Goods, error) {
	goods, err := model.GetGoods(in.Id)
	if err != nil {
		return nil, err
	}
	return &api.Goods{
		Id:           goods.Id,
		Name:         goods.Name,
		Price:        goods.Price.String(),
		Stock:        goods.Stock,
		Photo:        goods.Photo,
		Content:      goods.Content,
		BelongUserId: goods.BelongUsernameId,
	}, nil

}
