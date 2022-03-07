package model

import (
	"gorm.io/gorm"
	"time"
)

type StateFullEntity struct {
	gorm.Model
	Id         string `gorm:"primary_key type:varchar(255);not null'"`
	CreateTime time.Time
	UpdateTime time.Time
	State      string
}
