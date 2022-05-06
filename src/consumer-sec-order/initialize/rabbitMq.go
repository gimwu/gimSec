package initialize

import (
	"gimSec/basic/global"
	"gimSec/basic/logging"
	rabbit_mq "gimSec/basic/rabbit-mq"
	"gimSec/src/consumer-sec-order/server"
)

func Amqp() {
	connectionStr := "amqp://gimmick:123456@localhost:5672/cmm"
	global.CH = &rabbit_mq.MessageClient{}
	err := global.CH.NewConnection(connectionStr)
	if err != nil {
		logging.Error(err)
		return
	}

	err = global.CH.SubscribeToQueue("order", server.AddSecOrderByMq)
	if err != nil {
		logging.Error(err)
		return
	}
}
