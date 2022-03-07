package model

import (
	"gimSec/basic/model"
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
