package initialize

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"gimSec/basic/logging"
	"gimSec/src/provider-user/server"
)

func Dubbo() {
	config.SetProviderService(&server.UserProvider{})
	path := "./src/provider-user/provider-user.yaml"
	err := config.Load(config.WithPath(path))
	if err != nil {
		logging.Error(err.Error())
	}
}
