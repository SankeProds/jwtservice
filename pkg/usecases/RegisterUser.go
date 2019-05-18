package usecases

import (
	"fmt"
	"log"
)

// Import interface
type RegisterUserUC interface {
	RegisterUser(id string, data, authData interface{}, authMethod string) error
}

// Type remains private
type registerUserUC struct {
	repo UserRepo
}

// Create function
func NewRegisterUserUC(repo UserRepo) *registerUserUC {
	return &registerUserUC{
		repo: repo,
	}
}

// Interface implementation
func (u *registerUserUC) RegisterUser(id string, data, authData interface{}, authMethod string) error {
	log.Printf("Registering: [%s]", id)
	user, err := u.repo.FindById(id)
	if err != nil {
		return fmt.Errorf("Internal Error")
	} else if user != nil {
		return fmt.Errorf("user [%s] already exists", id)
	}
	user = NewAuthUser(id, data, authMethod, authData)
	if u.repo.Store(user) {
		return nil
	}
	return fmt.Errorf("unexpected error registering [%s]", id)
}
