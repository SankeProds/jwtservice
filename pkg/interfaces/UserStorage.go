package interfaces

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/SankeProds/jwtservice/pkg/domain"
)

/* Finds and Stores Users */

type StringStorage interface {
	Get(key string) (string, error)
	Save(key, val string) error
}

type UserStorage struct {
	storage StringStorage
}

// public create function
func NewUserStorage(storage StringStorage) *UserStorage {
	return &UserStorage{
		storage: storage,
	}
}

func (ur *UserStorage) FindByName(name string) *domain.User {
	strRep, err := ur.storage.Get(name)
	if err != nil {
		log.Printf("Error getting user: %+v", err)
		return nil
	}
	dec := json.NewDecoder(strings.NewReader(strRep))
	var user domain.User
	if err = dec.Decode(&user); err != nil {
		log.Printf("Error getting user: %+v", err)
		return nil
	}
	return &user
}

func (ur *UserStorage) Store(u *domain.User) bool {
	b, err := json.Marshal(u)
	if err != nil {
		log.Printf("Error storing user: %+v", err)
		return false
	}
	err = ur.storage.Save(u.Id, string(b))
	if err != nil {
		log.Printf("Error storing user: %+v", err)
	}
	return err == nil
}
