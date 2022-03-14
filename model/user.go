package model

import (
	"gimSec/basic/model"
)

type User struct {
	model.StateFullEntity
	Name      string `gorm:"type:varchar(255);not null"`
	Telephone string `gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
}
