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
