package model

import (
	"gorm.io/gorm"
)

type StateFullEntity struct {
	gorm.Model
	Id int64 `gorm:"primarykey type:varchar(255);not null'"`
}
