package server

import (
	"gimSec/basic/utils"
	"gimSec/model"
)

func CheckAdmin(admin *model.Admin) (bool, error) {
	return model.CheckAdmin(admin.Username)
}

func AddAdmin(admin *model.Admin) error {
	admin.Id = utils.SnowFlake.NextVal()
	return model.AddAdmin(admin)
}

func AdminLogin(admin *model.Admin) error {
	return model.AdminLogin(admin)
}
