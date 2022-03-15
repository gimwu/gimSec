package model

import (
	"gorm.io/gorm"
)

type StateFullEntity struct {
	gorm.Model
	Id string `gorm:"primarykey type:varchar(255);not null'"`
}
