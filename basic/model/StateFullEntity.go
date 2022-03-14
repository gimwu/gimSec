package model

import (
	"gorm.io/gorm"
)

type StateFullEntity struct {
	gorm.Model
	Id    string `gorm:"primary_key type:varchar(255);not null'"`
	State string
}
