package initialize

import (
	"gimSec/basic/logging"
	router2 "gimSec/src/provider-goods/router"
	"net/http"
)

func Router() {
	logging.Info("router Init")
	router := router2.InitRouter()
	logging.Info("Server Init")
	server := &http.Server{
		Addr:    ":8083",
		Handler: router,
	}

	logging.Info("server Listen")
	err := server.ListenAndServe()
	if err != nil {
		logging.Error("server Listen error :", err)
		return
	}
}
