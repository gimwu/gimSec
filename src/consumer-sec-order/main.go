package main

import "gimSec/src/consumer-sec-order/initialize"

func main() {
	initialize.Dubbo()
	initialize.Redis()
	initialize.Gorm()
	initialize.Router()
}
