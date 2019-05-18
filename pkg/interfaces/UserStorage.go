package interfaces

import (
	"log"

	"github.com/SankeProds/jwtservice/pkg/usecases"
)

/* Finds and Stores Users */

type AuthUserStorage interface {
	Get(id string) (*usecases.AuthUser, error)
	Save(*usecases.AuthUser) error
}

type UserStorage struct {
	storage AuthUserStorage
}

// public create function
func NewUserStorage(storage AuthUserStorage) *UserStorage {
	return &UserStorage{
		storage: storage,
	}
}

func (ur *UserStorage) FindById(id string) (*usecases.AuthUser, error) {
	authUser, err := ur.storage.Get(id)
	if err != nil {
		log.Printf("Error getting user: %+v", err)
		return nil, err
	}
	return authUser, nil
}

func (ur *UserStorage) Store(authUser *usecases.AuthUser) bool {
	err := ur.storage.Save(authUser)
	if err != nil {
		log.Printf("Error storing user: %+v", err)
	}
	return err == nil
}
