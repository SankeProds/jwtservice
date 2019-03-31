package domain

/* Simple user representation */

type User struct {
	Id       string
	Password string
}

func NewUser(id, password string) *User {
	return &User{
		Id:       id,
		Password: password,
	}
}

func (u User) CheckPassword(pass string) bool {
	return u.Password == pass
}
