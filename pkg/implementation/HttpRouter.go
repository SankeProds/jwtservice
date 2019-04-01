package implementation

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewEndpoint(path string, handler func(http.ResponseWriter, *http.Request), methods []string) *endpoint {
	return &endpoint{
		path:    path,
		handler: handler,
		methods: methods,
	}
}

type endpoint struct {
	path    string
	methods []string
	handler func(http.ResponseWriter, *http.Request)
}

type HttpApp interface {
	GetEndpoints() []*endpoint
}

type httpRouter struct {
	muxR *mux.Router
}

func NewHttpRouter() *httpRouter {
	return &httpRouter{
		muxR: mux.NewRouter(),
	}
}

func (router *httpRouter) RegisterApp(app HttpApp) {
	endpoints := app.GetEndpoints()
	for _, endpoint := range endpoints {
		router.muxR.HandleFunc(endpoint.path, endpoint.handler).Methods(endpoint.methods...)
	}
}
