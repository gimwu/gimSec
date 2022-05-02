package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"gimSec/api"
	"gimSec/basic/logging"
)

func main() {
	path := "./src/provider-order/provider-user.yaml"
	err := config.Load(config.WithPath(path))
	if err != nil {
		panic(err)
	}
	logging.Info("start to test triple unary context attachment transport")
	req := &api.IdMessage{
		Id: "laurence",
	}
}
