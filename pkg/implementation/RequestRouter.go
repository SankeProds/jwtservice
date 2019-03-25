package implementation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type App interface {
	RegisterHandlers(*mux.Router)
}

type RequestRouter struct {
	muxRouter *mux.Router
}

func (rr *RequestRouter) AddApp(app App) {
	app.RegisterHandlers(rr.muxRouter)
}

func (rr *RequestRouter) Init() {
	rr.muxRouter = mux.NewRouter()
}

func (rr *RequestRouter) GetHttpHandler() http.Handler {
	return rr.muxRouter
}
