package usecases

import "github.com/SankeProds/jwtservice/pkg/domain"

type UserRepo interface {
	FindByName(string) *domain.User
	Store(*domain.User) bool
}
