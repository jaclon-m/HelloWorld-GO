package main

import (
	"httpserver-demo/cfg"
	"httpserver-demo/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	var serverConfig = cfg.SetUp()
	mux := http.DefaultServeMux
	mux.Handle("/", handler.DefaultDispatchHandler)
	serv := &http.Server{
		Addr:              serverConfig.Address,
		ReadHeaderTimeout: 1000 * time.Millisecond,
		IdleTimeout:       1800 * 1000 * time.Millisecond,
		ReadTimeout:       1000 * time.Millisecond,
		Handler:           mux,
	}
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("start server failed!")
	}
}
