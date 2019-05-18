package usecases

import "github.com/SankeProds/jwtservice/pkg/domain"

type AuthUser struct {
	user       *domain.User
	authMethod string
	authData   interface{}
}

// repo interface to get users
type UserRepo interface {
	FindById(string) (*AuthUser, error)
	Store(*AuthUser) bool
}

func NewAuthUser(id string, data interface{}, authMethod string, authData interface{}) *AuthUser {
	return &AuthUser{
		user:       domain.NewUser(id, data),
		authMethod: authMethod,
		authData:   authData,
	}
}

func (user *AuthUser) GetId() string {
	return user.user.Id
}

func (user *AuthUser) GetData() interface{} {
	return user.user.Data
}

func (user *AuthUser) GetAuthMethod() string {
	return user.authMethod
}

func (user *AuthUser) GetAuthData() interface{} {
	return user.authData
}
