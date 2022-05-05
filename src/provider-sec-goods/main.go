package main

import "gimSec/src/provider-sec-goods/initialize"

func main() {
	initialize.Dubbo()
	initialize.Gorm()
	initialize.Redis()
	initialize.Router()
}
