package main

import "gimSec/src/consumer-admin/initialize"

func main() {
	initialize.Dubbo()
	initialize.Redis()
	initialize.Gorm()
	initialize.Router()
}
