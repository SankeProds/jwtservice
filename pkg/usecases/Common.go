package usecases

import "github.com/SankeProds/jwtservice/pkg/domain"

type AuthUser struct {
	user     *domain.User
	authData interface{}
}

// repo interface to get users
type UserRepo interface {
	FindById(string) (*AuthUser, error)
	Store(*AuthUser) bool
}

func NewAuthUser(id string, data, authData interface{}) *AuthUser {
	return &AuthUser{
		user:     domain.NewUser(id, data),
		authData: authData,
	}
}

func (user *AuthUser) GetId() string {
	return user.user.Id
}

func (user *AuthUser) GetData() interface{} {
	return user.user.Data
}

func (user *AuthUser) GetAuthData() interface{} {
	return user.authData
}

// Validator interface
type Authenticator interface {
	Validate(authData interface{}) error
	Authenticate(authData, loginData interface{}) error
}
