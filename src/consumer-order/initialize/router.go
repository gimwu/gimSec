package initialize

import (
	"gimSec/basic/logging"
	router2 "gimSec/src/consumer-order/router"
	"net/http"
)

func Router() {
	logging.Info("router Init")
	router := router2.InitRouter()
	logging.Info("Server Init")
	server := &http.Server{
		Addr:    ":8070",
		Handler: router,
	}

	logging.Info("server Listen")
	err := server.ListenAndServe()
	if err != nil {
		logging.Error("server Listen error :", err)
		return
	}
}
