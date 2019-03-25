package usecases

import (
	"errors"

	"github.com/SankeProds/jwtservice/pkg/domain"
)

type SessionUsecase interface {
	Login(name, password string) (string, error)
}

type sessionUsecase struct {
	repo domain.Repo
}

func NewSessionUsecase(repo domain.Repo) *sessionUsecase {
	return &sessionUsecase{
		repo: repo,
	}
}

func (u *sessionUsecase) Login(name, password string) (string, error) {
	user := u.repo.FindByName(name)
	if err := user == nil || !user.CheckPassword(password); err {
		return "", errors.New("Bad user/pass")
	}
	return getToken(user), nil
}

func getToken(u *domain.User) string {
	return "ELTOKEN"
}
