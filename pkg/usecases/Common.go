package usecases

import "github.com/SankeProds/jwtservice/pkg/domain"

// repo interface to get users
type UserRepo interface {
	FindByName(string) *domain.User
	Store(*domain.User) bool
}
