package main

import (
	router2 "gimSec/router"
	"net/http"
)

func main() {

	router := router2.InitRouter()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
