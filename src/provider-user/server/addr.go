package server

import (
	"gimSec/basic/utils"
	"gimSec/src/provider-user/model"
)

func AddAddr(addr *model.Addr) error {
	addr.Id = utils.SnowFlake.NextVal()
	return model.AddAddr(addr)
}

func GetAddr(id string) (*model.Addr, error) {
	return model.GetAddr(id)
}

func EditAddr(addr *model.Addr) error {
	return model.EditAddr(addr)
}

func DeleteAddr(id string) (*model.Addr, error) {
	addr, err := model.GetAddr(id)
	if err != nil {
		return nil, err
	}
	return model.DeleteAddr(addr)
}

func QueryAddrList(params map[string]string) ([]*model.Addr, error) {
	return model.QueryAddrList(params)
}
