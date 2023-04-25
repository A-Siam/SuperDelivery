package controller

import (
	"log"
	"net/http"
	"time"
)

var server *http.Server
var rootMux *http.ServeMux

func CreateServer(host string, port string, readTimeout time.Duration) *http.Server {
	if server == nil {
		rootMux = http.NewServeMux()
		server = &http.Server{
			Addr:        host + ":" + port,
			ReadTimeout: readTimeout * time.Second,
			Handler:     rootMux,
		}
	}
	return server
}

func GetRootMux() *http.ServeMux {
	if rootMux == nil {
		rootMux = http.NewServeMux()
	}
	return rootMux
}

func Start(server *http.Server) error {
	log.Default().Println("start listening on:", server.Addr)
	err := server.ListenAndServe()
	return err
}
