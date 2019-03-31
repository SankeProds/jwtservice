package implementation

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

/* Validates and extract data from http request and calls the User use cases */

type userApp struct {
	userHandler usecases.UserUsecase
}

// Public constructor
func NewUserApp(uc usecases.UserUsecase) *userApp {
	return &userApp{
		userHandler: uc,
	}
}

// Uses the input mux.Router to register where and how this apps expects its calls
func (app *userApp) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/user", app.registerUserHandler).Methods("POST")
}

func (app *userApp) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get and Check for params
	user, password, err := getNameAndPasswordFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
		return
	}
	// Make the call
	if err := app.userHandler.RegisterUser(user, password); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}
}
