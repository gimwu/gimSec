package server

//
//import (
//	"gimSec/basic/e"
//	"gimSec/basic/utils"
//	"gimSec/model"
//	"github.com/shopspring/decimal"
//)
//
//func AddGoodsOrderItem(goods *model.Goods, goodsNum int64, user *model.User) (*model.GoodsOrderItem, error) {
//	isExist, err := model.CheckGoodsStock(goods)
//	if err != nil {
//		return nil, err
//	}
//
//	if !isExist {
//		return nil, e.Inventory_Shortage_Error
//	}
//	goods.Stock -= goodsNum
//
//	price := goods.Price.Mul(decimal.NewFromInt(int64(goodsNum)))
//	goodsOrderItem := model.GoodsOrderItem{
//		GoodId:      goods.Id,
//		UsernameId:  user.Id,
//		GoodsNum:    goodsNum,
//		Price:       price,
//		OrderStatus: model.ORDER_NOT_PAY,
//	}
//	goodsOrderItem.Id = utils.SnowFlake.NextVal()
//
//	err = model.AddGoodsOrderItem(&goodsOrderItem)
//	if err != nil {
//		return nil, err
//	}
//	err = model.EditGoods(goods)
//	if err != nil {
//		return nil, err
//	}
//	return &goodsOrderItem, nil
//}
//
//func DeleteGoodsOrderItem(id string) (*model.GoodsOrderItem, error) {
//	goodsOrderItem, err := model.GetGoodsOrderItem(id)
//	if err != nil {
//		return nil, err
//	}
//	return model.DeleteGoodsOrderItem(goodsOrderItem)
//}
//
//func GetGoodsOrderItem(id string) (*model.GoodsOrderItem, error) {
//	return model.GetGoodsOrderItem(id)
//}
//
//func QueryGoodsOrderItemPage(params interface{}, currentPage int, pageSize int) (map[string]interface{}, error) {
//	goodsOrderItemList, err := model.QueryGoodsOrderItemPage(params, currentPage, pageSize)
//	if err != nil {
//		return nil, err
//	}
//	count, err := model.QueryGoodsOrderItemCount(params)
//	if err != nil {
//		return nil, err
//	}
//	res := make(map[string]interface{})
//	res["list"] = &goodsOrderItemList
//	res["count"] = count
//	return res, nil
//}
