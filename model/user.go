package model

import (
	"gimSec/basic/model"
	"time"
)

//User 用户表
type User struct {
	//继承父类
	model.StateFullEntity

	//email 邮箱
	Email string

	//nickname 昵称
	Nickname string

	//password 密码
	Password string

	//avatar 头像
	Avatar string

	//phone 电话
	Phone string

	//最后一次登录时间
	LastLoginTime time.Time

	//是否是商家
	IsErchant bool
}
