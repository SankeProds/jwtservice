package implementation

import (
	"encoding/json"
	"fmt"
)

type authenticator struct{}

type AuthData struct {
	AuthMethod string
	Password   string
}

func getData(data interface{}) (*AuthData, error) {
	asJson, err := json.Marshal(data.(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	var passData AuthData
	if err = json.Unmarshal(asJson, &passData); err != nil {
		return nil, err
	}
	return &passData, nil
}

func NewAuthenticator() *authenticator {
	return &authenticator{}
}

func (auth *authenticator) Validate(authData interface{}) error {
	data, err := getData(authData)
	if err != nil {
		return err
	}
	if data.AuthMethod == "password" {
		return validatePasswordData(data)
	}
	return fmt.Errorf("Unsuported auth method: [%+v]", data.AuthMethod)
}

func (auth *authenticator) Authenticate(authData, loginData interface{}) error {
	return nil
}

func validatePasswordData(passwordData *AuthData) error {
	if passwordData.Password == "" {
		return fmt.Errorf("No password provided")
	}
	return nil
}
