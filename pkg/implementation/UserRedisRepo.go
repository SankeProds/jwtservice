package implementation

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/go-redis/redis"

	"github.com/SankeProds/jwtservice/pkg/domain"
)

type UserRedisRepo struct {
	client *redis.Client
}

func NewUserRedisRepo() *UserRedisRepo {
	return &UserRedisRepo{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:1234", // use default Addr
			Password: "",               // no password set
			DB:       0,                // use default DB
		}),
	}
}

func (ur *UserRedisRepo) FindByName(name string) *domain.User {
	b, err := ur.client.Get(name).Bytes()
	if err != nil {
		log.Printf("Error getting user: %+v", err)
		return nil
	}
	dec := json.NewDecoder(strings.NewReader(string(b)))
	var user domain.User
	if err = dec.Decode(&user); err != nil {
		log.Printf("Error getting user: %+v", err)
		return nil
	}
	return &user
}

func (ur *UserRedisRepo) Store(u *domain.User) bool {
	b, err := json.Marshal(u)
	if err != nil {
		log.Printf("Error storing user: %+v", err)
		return false
	}
	err = ur.client.Set(u.Name, b, 0).Err()
	if err != nil {
		log.Printf("Error storing user: %+v", err)
	}
	return err == nil
}
