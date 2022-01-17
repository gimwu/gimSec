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
	Username string

	//nickname 昵称
	Nickname string

	//password 密码
	Password string

	//avatar 头像
	Avatar string

	//level 级别
	Level string

	//lastLoginTime 最后一次登录时间
	LastLoginTime time.Time
}
