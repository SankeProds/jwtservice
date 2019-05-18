package implementation

import (
	"encoding/json"
	"net/http"
)

type RequestData struct {
	Id          string
	Data        interface{}
	LoginMethod string
	AuthData    interface{}
	LoginData   interface{}
}

func getDataFromRequest(r *http.Request) (RequestData, error) {
	decoder := json.NewDecoder(r.Body)
	var data RequestData
	err := decoder.Decode(&data)
	return data, err
}
