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
	authenticator  Authenticator
}

// Create function
func NewLoginUC(repo UserRepo, tokenImpl TokenGenerator, authenticator Authenticator) *loginUC {
	return &loginUC{
		repo:           repo,
		tokenGenerator: tokenImpl,
		authenticator:  authenticator,
	}
}

// interface implementation
func (self *loginUC) Login(id string, loginMethod string, loginData interface{}) (string, error) {
	log.Printf("Loging in user: [%s]", id)
	user, err := self.repo.FindById(id)
	if fail := err != nil || user == nil; fail {
		log.Printf("bad user/pass: %+v", err)
		return "", errors.New("bad user/pass")
	}
	if err = self.authenticator.Authenticate(user.GetAuthData(), loginData); err != nil {
		log.Printf("bad user/pass: %+v", err)
		return "", errors.New("bad user/pass")
	}
	token, err := self.tokenGenerator.GetToken(user.GetId())
	if err != nil {
		return "", fmt.Errorf("err getting token: %+v", err)
	}
	return token, nil
}
