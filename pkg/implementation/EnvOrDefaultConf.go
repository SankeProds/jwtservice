package implementation

import (
	"log"
	"os"
	"strconv"
)

/* Simple configuration object. Could implement differenc services
   configuration necesities */

var DEFAULT = map[string]interface{}{
	"REDIS_ADDR":       "localhost:6379",
	"REDIS_PASSWORD":   "",
	"REDIS_DB":         0,
	"SERVICE_HOST":     "localhost",
	"SERVICE_PORT":     1234,
	"JWT_SIGNING_KEY":  "HolaEnfermera",
	"POSTGRES_CONNSTR": "postgres://postgres:123123@192.168.0.19/postgres?sslmode=disable",
}

type EnvOrDefaultConf struct{}

func NewEnvOrDefaultConf() *EnvOrDefaultConf {
	return &EnvOrDefaultConf{}
}

func (c *EnvOrDefaultConf) GetConnStr() string {
	return lookStringInEnv("POSTGRES_CONNSTR")
}

func (c *EnvOrDefaultConf) GetSigningKey() string {
	return lookStringInEnv("JWT_SIGNING_KEY")
}

func (c *EnvOrDefaultConf) GetPort() int {
	return lookIntInEnv("SERVICE_PORT")
}

func (c *EnvOrDefaultConf) GetHost() string {
	return lookStringInEnv("SERVICE_HOST")
}

func (c *EnvOrDefaultConf) GetRedisAddr() string {
	return lookStringInEnv("REDIS_ADDR")
}

func (c *EnvOrDefaultConf) GetRedisPassword() string {
	return lookStringInEnv("REDIS_PASSWORD")
}

func (c *EnvOrDefaultConf) GetRedisDB() int {
	return lookIntInEnv("REDIS_DB")
}

func lookStringInEnv(patern string) string {
	if val, ok := os.LookupEnv(patern); ok {
		return val
	}
	log.Printf("%s: Not found in enviroment. Using [%+v]", patern, DEFAULT[patern])
	return DEFAULT[patern].(string)
}

func lookIntInEnv(patern string) int {
	if val, ok := os.LookupEnv(patern); ok {
		if valInt, ok := strconv.Atoi(val); ok == nil {
			return valInt
		}
	}
	log.Printf("%s: Not found in enviroment. Using [%+v]", patern, DEFAULT[patern])
	return DEFAULT[patern].(int)
}
