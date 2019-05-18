package implementation

import (
	"fmt"
	"net/http"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

/* Validates and extract data from http request and calls the Session use cases */

type LoginApp struct {
	loginUC usecases.LoginUC
}

// Public constructor
func NewLoginApp(loginUC usecases.LoginUC) *LoginApp {
	return &LoginApp{
		loginUC: loginUC,
	}
}

func (app *LoginApp) GetEndpoints() []*endpoint {
	e := NewEndpoint("/login", app.loginHandler, []string{"POST"})
	return []*endpoint{e}
}

func (app *LoginApp) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Get and Check for params
	data, err := getDataFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
		return
	}
	// Make the call
	if token, err := app.loginUC.Login(data.Id, data.LoginMethod, data.LoginData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, token)
	}
}
