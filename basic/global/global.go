package global

import (
	rabbit_mq "gimSec/basic/rabbit-mq"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	REDIS *redis.Client
	DB    *gorm.DB
	CH    rabbit_mq.IMessageClient
)
