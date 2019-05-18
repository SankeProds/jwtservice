package domain

/* Simple user representation */

type User struct {
	Id   string
	Data interface{}
}

func NewUser(id string, data interface{}) *User {
	return &User{
		Id:   id,
		Data: data,
	}
}
