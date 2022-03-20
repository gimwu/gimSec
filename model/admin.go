package model

import (
	"gimSec/basic/model"
	"gorm.io/gorm"
	"time"
)

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

func AdminLogin(admin *Admin) error {
	err := db.Where("Username = ? and Password = ?", admin.Username, admin.Password).First(&admin).Error
	if err != nil {
		return err
	}
	return nil
}
