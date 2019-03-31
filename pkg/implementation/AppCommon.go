package implementation

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

/* Interfaces and function that could be util for several apps */

type App interface {
	RegisterHandlers(*mux.Router)
}

func getNameAndPasswordFromRequest(r *http.Request) (string, string, error) {
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
