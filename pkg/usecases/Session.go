package usecases

import (
	"errors"
	"fmt"
	"log"
)

type JwtGenerator interface {
	GetToken(string) (string, error)
}

// Import interface
type SessionUsecase interface {
	Login(name, password string) (string, error)
}

// Type remains private
type sessionUsecase struct {
	repo   UserRepo
	jwtGen JwtGenerator
}

// Create function
func NewSessionUsecase(repo UserRepo, jwtImplementation JwtGenerator) *sessionUsecase {
	return &sessionUsecase{
		repo:   repo,
		jwtGen: jwtImplementation,
	}
}

// interface implementation
func (u *sessionUsecase) Login(name, password string) (string, error) {
	log.Printf("Loging in user: %s", name)
	user := u.repo.FindByName(name)
	if err := user == nil || !user.CheckPassword(password); err {
		return "", errors.New("Bad user/pass")
	}
	token, err := u.jwtGen.GetToken(user.Name)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Err getting token %+v", err))
	}
	return token, nil
}
