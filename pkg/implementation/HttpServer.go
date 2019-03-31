package implementation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* Handles the server instrumentation */

type ServerConf interface {
	GetHost() string
	GetPort() int
}

type HttpServer struct {
	host      string
	port      int
	reqRouter *mux.Router
}

func (server *HttpServer) Init(conf ServerConf, reqRouter *mux.Router) {
	server.host = conf.GetHost()
	server.port = conf.GetPort()
	server.reqRouter = reqRouter
}

func (server *HttpServer) Start() {
	http.Handle("/", server.reqRouter)
	serverUrl := fmt.Sprintf("%s:%d", server.host, server.port)
	log.Printf(fmt.Sprintf("Now listening on: %s", serverUrl))
	// For security reasons, You should change this to: ListenAndServeTLS
	// see https://golang.org/pkg/net/http/#ListenAndServeTLS form more info
	log.Fatal(http.ListenAndServe(serverUrl, nil))
}
