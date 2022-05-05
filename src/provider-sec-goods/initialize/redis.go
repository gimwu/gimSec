package initialize

import (
	"gimSec/basic/global"
	"github.com/go-redis/redis/v8"
)

func Redis() {
	global.REDIS = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default USER_DB
	})
}
