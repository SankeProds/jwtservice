package implementation

import (
	"fmt"
	"net/http"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

/* Validates and extract data from http request and calls the Session use cases */

type AuthUserApp struct {
	authUsers usecases.AuthUserUC
}

// Public constructor
func NewAuthUserApp(authUsers usecases.AuthUserUC) *AuthUserApp {
	return &AuthUserApp{
		authUsers: authUsers,
	}
}

func (app *AuthUserApp) GetEndpoints() []*endpoint {
	e := NewEndpoint("/auth", app.loginHandler, []string{"POST"})
	return []*endpoint{e}
}

func (app *AuthUserApp) loginHandler(w http.ResponseWriter, r *http.Request) {
	// Get and Check for params
	data, err := getDataFromRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
		return
	}
	// Make the call
	if token, err := app.authUsers.Login(data.Id, data.LoginMethod, data.LoginData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%+v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, token)
	}
}
