package main

import "gimSec/src/consumer-order/initialize"

func main() {
	initialize.Dubbo()
	initialize.Gorm()
	initialize.Router()
}
