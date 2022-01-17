package dao

import (
	"gimSec/basic/utils"
	"gimSec/model"
)

var gormConnection = utils.GormMysqlDatabase

//AddAdmin 增
func AddAdmin(admin *model.Admin) (*model.Admin, error) {
	err := gormConnection.Create(&admin).Error
	if err != nil {
		return nil, err
	}
	return admin, nil
}

//RemoveAdmin 删
func RemoveAdmin(id string) (*model.Admin, error) {
	var admin model.Admin
	err := gormConnection.Model(&admin).Where("id = ? And status = ?", id, "ENABLE").Update("status", "DefaultGoTemplateProperty").Error
	if err != nil {
		return nil, err
	}
	return &admin, err
}

//EditAdmin 改
func EditAdmin(id string, data interface{}) (*model.Admin, error) {
	var admin model.Admin
	err := gormConnection.Model(&admin).Where("id = ? And status = ?", id, "ENABLE").Updates(data).Error
	if err != nil {
		return nil, err
	}
	return &admin, err
}

//GetAdmin 查
func GetAdmin(id string) *model.Admin {
	var admin model.Admin
	err := gormConnection.Model(&admin).Where("id = ? And status = ?", id, "ENABLE").Error
	if err != nil {
		panic(err)
	}
	return &admin
}
