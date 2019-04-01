package implementation

import (
	"github.com/go-redis/redis"
)

type RedisConf interface {
	GetRedisAddr() string
	GetRedisPassword() string
	GetRedisDB() int
}

type redisStringStorage struct {
	client *redis.Client
}

func NewRedisStringStorage(conf RedisConf) *redisStringStorage {
	return &redisStringStorage{
		client: redis.NewClient(&redis.Options{
			Addr:     conf.GetRedisAddr(),
			Password: conf.GetRedisPassword(),
			DB:       conf.GetRedisDB(),
		}),
	}
}

func (rss *redisStringStorage) Get(key string) (string, error) {
	b, err := rss.client.Get(key).Bytes()
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (rss *redisStringStorage) Save(key, val string) error {
	return rss.client.Set(key, val, 0).Err()
}
