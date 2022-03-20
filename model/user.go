package model

import (
	"gimSec/basic/model"
	"gorm.io/gorm"
	"time"
)

type User struct {
	model.StateFullEntity
	Name          string `gorm:"type:varchar(255);not null"`
	Telephone     string `gorm:"type:varchar(255);not null;unique"`
	Password      string `gorm:"type:varchar(255);not null"`
	LastTimeLogin time.Time
}

func CheckUser(telephone string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Telephone: telephone}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if len(user.Id) > 0 {
		return true, nil
	}

	return false, nil
}

func AddUser(user *User) error {
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id string) (*User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func EditUser(user *User) error {
	err := db.Updates(user).Error
	if err != nil {
		return err
	}
	return nil
}

func Login(user *User) error {
	err := db.Where("Telephone = ? and Password =?", user.Telephone, user.Password).First(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func QueryUserPage(params interface{}, currentPage int, pageSize int) ([]*User, error) {
	var UserList []*User
	err := db.Model(&User{}).Offset((currentPage - 1) * pageSize).Limit(pageSize).Find(&UserList).Error
	if err != nil {
		return nil, err
	}
	return UserList, nil
}

func QueryUserCount(params interface{}) (int64, error) {
	var count int64
	err := db.Model(&User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
