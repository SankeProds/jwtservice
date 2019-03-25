package implementation

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

type userApp struct {
	userHandler usecases.UserUsecase
}

func (app *userApp) RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/user", app.registerUserHandler).Methods("POST")
}

func (app *userApp) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get and Check for params
	user, password, err := getParamsFromRequest(r)
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

func NewUserApp(uc usecases.UserUsecase) *userApp {
	return &userApp{
		userHandler: uc,
	}
}

func getParamsFromRequest(r *http.Request) (string, string, error) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		return "", "", errors.New("missing param: name")
	}
	password := query.Get("password")
	if password == "" {
		return "", "", errors.New("missing param: password")
	}
	return name, password, nil
}
