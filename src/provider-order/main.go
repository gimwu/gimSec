package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"gimSec/api"
	"gimSec/basic/logging"
)

var grpcGreeterImpl = new(api.UserServiceClientImpl)

func init() {
	config.SetConsumerService(grpcGreeterImpl)
}

func main() {
	path := "./src/provider-order/provider-order.yaml"
	err := config.Load(config.WithPath(path))
	if err != nil {
		panic(err)
	}
	logging.Info("start to test triple unary context attachment transport")
	req := &api.IdMessage{
		Id: "154578522755563520",
	}
	reply, err := grpcGreeterImpl.GetUserById(context.Background(), req)
	if err != nil {
		logging.Error(err)
	}
	logging.Info("client response result1111111111: %v\n", reply.Telephone)

}
