package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	REDIS *redis.Client
	DB    *gorm.DB
)
