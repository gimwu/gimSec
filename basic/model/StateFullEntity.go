package model

import "time"

type StateFullEntity struct {
	id         string
	createTime time.Time
	updateTime time.Time
}
