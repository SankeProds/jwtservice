package implementation

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

type sessionApp struct {
	sessionHandler usecases.SessionUsecase
}

func (app *sessionApp) RegisterHandlers(r *mux.Router) {
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
		fmt.Fprintf(w, token)
	}
}

func NewSessionApp(sessionCases usecases.SessionUsecase) *sessionApp {
	return &sessionApp{
		sessionHandler: sessionCases,
	}
}
