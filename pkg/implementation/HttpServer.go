package implementation

import (
	"fmt"
	"log"
	"net/http"
)

/* Handles the server instrumentation */

type ServerConf interface {
	GetHost() string
	GetPort() int
}

type HttpServer struct {
	host   string
	port   int
	router *httpRouter
}

func (server *HttpServer) Init(conf ServerConf, router *httpRouter) {
	server.host = conf.GetHost()
	server.port = conf.GetPort()
	server.router = router
}

func (server *HttpServer) Start() {
	http.Handle("/", server.router.muxR)
	serverUrl := fmt.Sprintf("%s:%d", server.host, server.port)
	log.Println(fmt.Sprintf("Now listening on: %s", serverUrl))
	// For security reasons, You should change this to: ListenAndServeTLS
	// see https://golang.org/pkg/net/http/#ListenAndServeTLS form more info
	log.Fatal(http.ListenAndServe(serverUrl, nil))
}
