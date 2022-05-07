package server

import (
	"context"
	"gimSec/api"
)

type GoodsProvider struct {
	api.UnimplementedGoodsServiceServer
}

func (p *GoodsProvider) GetGoodsById(ctx context.Context, in *api.GoodsId) (*api.Goods, error) {
	goods, err := GetGoods(in.Id)
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

func (p *GoodsProvider) GetGoodsByIds(ctx context.Context, in *api.GoodsIds) (*api.Goodss, error) {
	var ids []string
	ids = in.Id
	goodsList, err := SelectGoodsByIds(ids)
	if err != nil {
		return nil, err
	}

	goodssList := make([]*api.Goods, 0)
	for _, goods := range goodsList {
		goods := &api.Goods{
			Id:           goods.Id,
			Name:         goods.Name,
			Price:        goods.Price.String(),
			Stock:        goods.Stock,
			Photo:        goods.Photo,
			Content:      goods.Content,
			BelongUserId: goods.BelongUsernameId,
		}
		goodssList = append(goodssList, goods)
	}

	goodss := &api.Goodss{
		Goods: goodssList,
	}
	return goodss, nil
}
