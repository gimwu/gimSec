package main

import (
	"gimSec/src/provider-user/initialize"
)

func main() {
	initialize.Redis()
	initialize.Dubbo()
	initialize.Gorm()
	initialize.Router()
}
