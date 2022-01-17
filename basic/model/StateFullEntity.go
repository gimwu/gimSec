package model

import "time"

type StateFullEntity struct {
	Id         string
	CreateTime time.Time
	UpdateTime time.Time
	State      string
}
