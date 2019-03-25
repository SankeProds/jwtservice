package domain

type Repo interface {
	FindByName(string) *User
	Store(*User) bool
}

type User struct {
	Name     string
	Password string
}

func NewUser(name, password string) *User {
	return &User{
		Name:     name,
		Password: password,
	}
}

func (u User) CheckPassword(pass string) bool {
	return u.Password == pass
}
