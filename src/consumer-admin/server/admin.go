package server

import (
	"gimSec/basic/utils"
	"gimSec/src/consumer-admin/model"
	"time"
)

func CheckAdmin(admin *model.Admin) (bool, error) {
	return model.CheckAdmin(admin.Username)
}

func AddAdmin(admin *model.Admin) error {
	admin.Id = utils.SnowFlake.NextVal()
	admin.LastLoginTime = time.Now()
	return model.AddAdmin(admin)
}

func GetAdmin(id string) (*model.Admin, error) {
	return model.GetAdmin(id)
}

func EditAdmin(admin *model.Admin) error {
	return model.EditAdmin(admin)
}

func DeleteAdmin(id string) (*model.Admin, error) {
	admin, err := model.GetAdmin(id)
	if err != nil {
		return nil, err
	}
	return model.DeleteAdmin(admin)
}

func QueryAdminPage(params interface{}, currentPage int, pageSize int) (map[string]interface{}, error) {
	adminList, err := model.QueryAdminPage(params, currentPage, pageSize)
	if err != nil {
		return nil, err
	}
	count, err := model.QueryAdminCount(params)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	res["list"] = &adminList
	res["count"] = count
	return res, nil
}

func AdminLogin(admin *model.Admin) error {
	return model.AdminLogin(admin)
}
