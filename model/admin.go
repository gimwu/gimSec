package model

import (
	"gimSec/basic/global"
	"gimSec/basic/model"
	"gorm.io/gorm"
	"time"
)

var db = global.ORDER_DB

//Admin 管理员表
type Admin struct {
	//继承父类
	model.StateFullEntity

	//username 账号
	Username string `gorm:"type:varchar(255);not null"`

	//nickname 昵称
	Nickname string `gorm:"type:varchar(255)"`

	//password 密码
	Password string `gorm:"type:varchar(255);not null"`

	//avatar 头像
	Avatar string `gorm:"type:varchar(255)"`

	//level 级别
	Level string `gorm:"type:varchar(255)"`

	//lastLoginTime 最后一次登录时间
	LastLoginTime time.Time
}

func CheckAdmin(username string) (bool, error) {
	var admin Admin
	err := db.Select("id").Where(Admin{Username: username}).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if len(admin.Id) > 1 {
		return true, nil
	}
	return false, nil
}

func AddAdmin(admin *Admin) error {
	err := db.Create(admin).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAdmin(id string) (*Admin, error) {
	var admin Admin
	err := db.Where("id = ?", id).First(&admin).Error
	if err != nil {
		return nil, err
	}
	return &admin, nil
}

func EditAdmin(admin *Admin) error {
	if err := db.Updates(admin).Error; err != nil {
		return err
	}
	return nil
}

func DeleteAdmin(admin *Admin) (*Admin, error) {
	if err := db.Delete(&admin).Error; err != nil {
		return nil, err
	}
	return admin, nil
}

func QueryAdminPage(params interface{}, currentPage int, pageSize int) ([]*Admin, error) {
	var adminList []*Admin
	err := db.Model(&Admin{}).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&adminList).Error
	if err != nil {
		return nil, err
	}
	return adminList, nil
}

func QueryAdminCount(params interface{}) (int64, error) {
	var count int64
	err := db.Model(&Admin{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil

}

func AdminLogin(admin *Admin) error {
	err := db.Where("Username = ? and Password = ?", admin.Username, admin.Password).First(&admin).Error
	if err != nil {
		return err
	}
	return nil
}
