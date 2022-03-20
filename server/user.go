package server

import (
	"gimSec/basic/utils"
	"gimSec/model"
)

func Check(user *model.User) (bool, error) {
	return model.CheckUser(user.Telephone)
}

func AddUser(user *model.User) error {
	user.Id = utils.SnowFlake.NextVal()
	return model.AddUser(user)
}

func GetUser(id string) (*model.User, error) {
	return model.GetUser(id)
}

func EditUser(user *model.User) error {
	return model.EditUser(user)
}

func Login(user *model.User) error {
	return model.Login(user)
}

func QueryUserPage(params interface{}, currentPage int, pageSize int) (map[string]interface{}, error) {
	userList, err := model.QueryUserPage(params, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QueryUserCount(params)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &userList
	res["count"] = count
	return res, nil
}
