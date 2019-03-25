package implementation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SankeProds/jwtservice/pkg/interfaces"
)

type HttpRequestRouter interface {
	GetHttpHandler() http.Handler
}

type HttpServer struct {
	host      string
	port      int
	reqRouter HttpRequestRouter
}

func (server *HttpServer) Init(conf interfaces.ConfInterface, reqRouter HttpRequestRouter) {
	server.host = conf.GetHost()
	server.port = conf.GetPort()
	server.reqRouter = reqRouter
}

func (server *HttpServer) Start() {
	http.Handle("/", server.reqRouter.GetHttpHandler())
	serverUrl := fmt.Sprintf("%s:%d", server.host, server.port)
	log.Printf(fmt.Sprintf("Now listening on: %s", serverUrl))
	log.Fatal(http.ListenAndServe(serverUrl, nil))
}
