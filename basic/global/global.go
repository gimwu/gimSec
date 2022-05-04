package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	USER_REDIS  *redis.Client
	GOODS_REDIS *redis.Client
	USER_DB     *gorm.DB
	GOODS_DB    *gorm.DB
)
