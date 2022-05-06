package main

import "gimSec/src/consumer-sec-order/initialize"

func main() {
	initialize.Dubbo()
	initialize.Redis()
	initialize.Amqp()
	initialize.Gorm()
	initialize.Router()
}
