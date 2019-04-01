package implementation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

/* Validates and extract data from http request and calls the Session use cases */

type sessionApp struct {
	sessionHandler usecases.SessionUsecase
}

// Public constructor
func NewSessionApp(sessionCases usecases.SessionUsecase) *sessionApp {
	return &sessionApp{
		sessionHandler: sessionCases,
	}
}

func (app *sessionApp) GetEndpoints() []*endpoint {
	e := NewEndpoint("/session", app.loginHandler, []string{"POST"})
	return []*endpoint{e}
}

// Uses the input mux.Router to register where and how this apps expects its calls
func (app *sessionApp) UseRouter(r *mux.Router) {
	log.Printf("login handler: [POST] /session")
	r.HandleFunc("/session", app.loginHandler).Methods("POST")
}

func (app *sessionApp) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Get and Check for params
	user, password, err := getNameAndPasswordFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
		return
	}
	// Make the call
	if token, err := app.sessionHandler.Login(user, password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, token)
	}
}
