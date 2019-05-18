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
type LoginUC interface {
	Login(id, loginMethod string, loginData interface{}) (string, error)
}

// Type remains private
type loginUC struct {
	repo           UserRepo
	tokenGenerator TokenGenerator
}

// Create function
func NewLoginUC(repo UserRepo, tokenImpl TokenGenerator) *loginUC {
	return &loginUC{
		repo:           repo,
		tokenGenerator: tokenImpl,
	}
}

// interface implementation
func (self *loginUC) Login(id string, loginMethod string, loginData interface{}) (string, error) {
	log.Printf("Loging in user: [%s]", id)
	user, err := self.repo.FindById(id)
	if err := err != nil || user == nil; err {
		return "", errors.New("bad user/pass")
	}
	token, err := self.tokenGenerator.GetToken(user.GetId())
	if err != nil {
		return "", fmt.Errorf("err getting token: %+v", err)
	}
	return token, nil
}
