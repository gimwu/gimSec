package main

import "gimSec/src/provider-goods/initialize"

func main() {
	initialize.Gorm()
	initialize.Router()
}
