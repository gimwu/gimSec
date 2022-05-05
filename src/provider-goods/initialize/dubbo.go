package initialize

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"gimSec/basic/logging"
	"gimSec/src/provider-goods/server"
)

func Dubbo() {
	config.SetProviderService(&server.GoodsProvider{})
	path := "./src/provider-goods/provider-goods.yaml"
	err := config.Load(config.WithPath(path))
	if err != nil {
		logging.Error(err.Error())
	}
}
