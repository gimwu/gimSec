package model

import (
	"gimSec/basic/global"
	"gimSec/basic/model"
	"gorm.io/gorm"
	"time"
)

var db = global.USER_DB

type User struct {
	model.StateFullEntity
	Name          string    `gorm:"type:varchar(255);not null" json:"name"`
	Telephone     string    `gorm:"type:varchar(255);not null;unique" json:"telephone"`
	Password      string    `gorm:"type:varchar(255);not null" json:"passWord"`
	UserType      string    `gorm:"type:varchar(255);not null" json:"userType"`
	LastTimeLogin time.Time `gorm:"type:datetime"json:"LastTimeLogin"`
}

func CheckUser(telephone string) (bool, error) {
	var user User
	err := global.USER_DB.Select("id").Where(User{Telephone: telephone}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if len(user.Id) > 0 {
		return true, nil
	}

	return false, nil
}

func AddUser(user *User) error {
	err := global.USER_DB.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id string) (*User, error) {
	var user User
	err := global.USER_DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func EditUser(user *User) error {
	err := global.USER_DB.Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User) (*User, error) {
	if err := global.USER_DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func Login(user *User) error {
	err := global.USER_DB.Where("Telephone = ? and Password =?", user.Telephone, user.Password).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryUserPage(params interface{}, currentPage int, pageSize int) ([]*User, error) {
	var UserList []*User
	err := global.USER_DB.Model(&User{}).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&UserList).Error
	if err != nil {
		return nil, err
	}
	return UserList, nil
}

func QueryUserCount(params interface{}) (int64, error) {
	var count int64
	err := global.USER_DB.Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
