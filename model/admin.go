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
	username string

	//nickname 昵称
	nickname string

	//password 密码
	password string

	//avatar 头像
	avatar string

	//level 级别
	level string

	//lastLoginTime 最后一次登录时间
	lastLoginTime time.Time
}
