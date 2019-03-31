package implementation

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

/* Interfaces and function that could be util for several apps */

type App interface {
	RegisterHandlers(*mux.Router)
}

func getNameAndPasswordFromRequest(r *http.Request) (string, string, error) {
	decoder := json.NewDecoder(r.Body)
	var data struct {
		User     string
		Password string
	}
	err := decoder.Decode(&data)
	if err != nil {
		return "", "", err
	}
	if data.User == "" {
		return "", "", errors.New("missing param: name")
	}
	if data.Password == "" {
		return "", "", errors.New("missing param: password")
	}
	return data.User, data.Password, nil
}
