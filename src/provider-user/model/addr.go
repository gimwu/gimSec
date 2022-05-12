package model

import (
	"gimSec/basic/global"
	"gimSec/basic/model"
)

type Addr struct {
	model.StateFullEntity
	UserId   string `gorm:"type:varchar(255);not null" json:"userId"`
	Addr     string `gorm:"type:varchar(255);not null" json:"addr"`
	UserName string `gorm:"type:varchar(255);not null" json:"username"`
}

func AddAddr(addr *Addr) error {
	err := global.DB.Create(&addr).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAddr(id string) (*Addr, error) {
	var addr Addr
	err := global.DB.Where("id = ?", id).First(&addr).Error
	if err != nil {
		return nil, err
	}
	return &addr, nil
}

func EditAddr(addr *Addr) error {
	err := global.DB.Updates(addr).Error
	return err
}

func DeleteAddr(addr *Addr) (*Addr, error) {
	if err := global.DB.Delete(&addr).Error; err != nil {
		return nil, err
	}
	return addr, nil
}

func QueryAddrList(params map[string]string) ([]*Addr, error) {
	var addrList []*Addr
	tx := global.DB.Model(&Addr{})

	if params["userId"] != "" {
		tx.Where("user_id = ?", params["userId"])
	}
	err := tx.Find(&addrList).Error
	if err != nil {
		return nil, err
	}
	return addrList, nil
}
