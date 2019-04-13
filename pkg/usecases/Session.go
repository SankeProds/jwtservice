package usecases

import (
	"errors"
	"fmt"
	"log"
)

type TokenGenerator interface {
	GetToken(string) (string, error)
}

// Import interface
type SessionUsecase interface {
	Login(name, password string) (string, error)
}

// Type remains private
type sessionUsecase struct {
	repo           UserRepo
	tokenGenerator TokenGenerator
}

// Create function
func NewSessionUsecase(repo UserRepo, tokenImpl TokenGenerator) *sessionUsecase {
	return &sessionUsecase{
		repo:           repo,
		tokenGenerator: tokenImpl,
	}
}

// interface implementation
func (self *sessionUsecase) Login(name, password string) (string, error) {
	log.Printf("Loging in user: %s", name)
	user := self.repo.FindByName(name)
	if err := user == nil || !user.CheckPassword(password); err {
		return "", errors.New("bad user/pass")
	}
	token, err := self.tokenGenerator.GetToken(user.Id)
	if err != nil {
		return "", fmt.Errorf("err getting token %+v", err)
	}
	return token, nil
}
