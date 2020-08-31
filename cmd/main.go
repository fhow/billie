package main

import (
	"github.com/billie/internal/config"
	"github.com/billie/internal/services/marstime/converter"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/convert", converter.Handler)
	configuration, err := config.InitConfig("billie")

	if configuration == nil || err != nil {
		log.Panic("config not loaded")
	}

	err = http.ListenAndServe(configuration.Server.Port, mux)
	if err != nil {
		log.Panic("server not started")
	}
}
