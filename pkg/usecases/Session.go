package usecases

import (
	"errors"

	"github.com/SankeProds/jwtservice/pkg/domain"
)

// Import interface
type SessionUsecase interface {
	Login(name, password string) (string, error)
}

// Type remains private
type sessionUsecase struct {
	repo UserRepo
	// TODO: jwt as a service
}

// Create function
func NewSessionUsecase(repo UserRepo) *sessionUsecase {
	return &sessionUsecase{
		repo: repo,
	}
}

// interface implementation
func (u *sessionUsecase) Login(name, password string) (string, error) {
	user := u.repo.FindByName(name)
	if err := user == nil || !user.CheckPassword(password); err {
		return "", errors.New("Bad user/pass")
	}
	return getToken(user), nil
}

func getToken(u *domain.User) string {
	// TODO use jwt service
	return "ELTOKEN"
}
