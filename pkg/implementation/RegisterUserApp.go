package implementation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

/* Validates and extract data from http request and calls the User use cases */

type registerUserApp struct {
	registerUserUC usecases.RegisterUserUC
}

// Public constructor
func NewRegisterUserApp(registerUserUC usecases.RegisterUserUC) *registerUserApp {
	return &registerUserApp{
		registerUserUC: registerUserUC,
	}
}

func (app *registerUserApp) GetEndpoints() []*endpoint {
	e := NewEndpoint("/user", app.registerUserHandler, []string{"POST"})
	return []*endpoint{e}
}

func (app *registerUserApp) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	// Get and Check for params
	data, err := getDataFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
		return
	}
	log.Printf("%+v", data.AuthMethod)
	if data.Id == "" || data.Data == nil || data.AuthData == nil || data.AuthMethod == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", "Missing required param")
		return
	}
	// Make the call
	if err := app.registerUserUC.RegisterUser(data.Id, data.Data, data.AuthData, data.AuthMethod); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")
	}
}
