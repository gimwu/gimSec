package main

import "gimSec/src/provider-goods/initialize"

func main() {
	initialize.Dubbo()
	initialize.Redis()
	initialize.Gorm()
	initialize.Router()
}
