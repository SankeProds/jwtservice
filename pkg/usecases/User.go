package usecases

import (
	"errors"
	"fmt"
	"log"

	"github.com/SankeProds/jwtservice/pkg/domain"
)

// Import interface
type UserUsecase interface {
	RegisterUser(name, password string) error
}

// Type remains private
type userUsecase struct {
	repo UserRepo
}

// Create function
func NewUserUsecase(repo UserRepo) *userUsecase {
	return &userUsecase{
		repo: repo,
	}
}

// Interface implementation
func (u *userUsecase) RegisterUser(name, password string) error {
	log.Printf("Registering: %s %s", name, password)
	user := u.repo.FindByName(name)
	if user != nil {
		return errors.New(fmt.Sprintf("User %s already exists!", name))
	}
	user = domain.NewUser(name, password)
	if u.repo.Store(user) {
		return nil
	}
	return errors.New(fmt.Sprintf("Unexpected error registering %s", name))
}
