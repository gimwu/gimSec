package initialize

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"gimSec/basic/logging"
	"gimSec/src/consumer-sec-order/server"
)

func Dubbo() {
	config.SetConsumerService(server.SecGoodsConsumer)
	path := "./src/consumer-sec-order/consumer-sec-order.yaml"
	err := config.Load(config.WithPath(path))
	if err != nil {
		logging.Error(err.Error())
	}
}
