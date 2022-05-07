package initialize

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"gimSec/basic/logging"
	"gimSec/src/consumer-order/server"
)

func Dubbo() {
	config.SetConsumerService(server.GoodsConsumer)
	path := "./src/consumer-order/consumer-order.yaml"
	err := config.Load(config.WithPath(path))
	if err != nil {
		logging.Error(err.Error())
	}
}
