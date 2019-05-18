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

func (auth *authenticator) Authenticate(authDataRaw, loginDataRaw interface{}) error {
	authData, err := getData(authDataRaw)
	loginData, err := getData(loginDataRaw)
	if err != nil {
		return fmt.Errorf("Error building login Data: %+v", err)
	}
	if authData.AuthMethod == "password" && authData.Password == loginData.Password {
		return nil
	}
	return fmt.Errorf("AuthFailed")
}

func validatePasswordData(passwordData *AuthData) error {
	if passwordData.Password == "" {
		return fmt.Errorf("No password provided")
	}
	return nil
}
