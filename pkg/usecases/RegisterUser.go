package usecases

import (
	"fmt"
	"log"
)

// Import interface
type RegisterUserUC interface {
	RegisterUser(id string, data, authData interface{}) error
}

// Type remains private
type registerUserUC struct {
	repo          UserRepo
	authenticator Authenticator
}

// Create function
func NewRegisterUserUC(repo UserRepo, authenticator Authenticator) *registerUserUC {
	return &registerUserUC{
		repo:          repo,
		authenticator: authenticator,
	}
}

// Interface implementation
func (u *registerUserUC) RegisterUser(id string, data, authData interface{}) error {
	log.Printf("Registering user: [%s]", id)
	user, err := u.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("Internal Error")
	} else if user != nil {
		return fmt.Errorf("user [%s] already exists", id)
	}
	if err := u.authenticator.Validate(authData); err != nil {
		return fmt.Errorf("Error validating registering data: %+v", err)
	}
	user = NewAuthUser(id, data, authData)
	if u.repo.Store(user) {
		return nil
	}
	return fmt.Errorf("unexpected error registering [%s]", id)
}
