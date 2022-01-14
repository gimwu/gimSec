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
	email string

	//nickname 昵称
	nickname string

	//password 密码
	password string

	//avatar 头像
	avatar string

	//phone 电话
	phone string

	//最后一次登录时间
	lastLoginTime time.Time

	//是否是商家
	isErchant bool
}
