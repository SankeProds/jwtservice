package implementation

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTGenerator struct {
	key string
}

type JWTConfiguration interface {
	GetSigningKey() string
}

func NewJWTGenerator(conf JWTConfiguration) *JWTGenerator {
	return &JWTGenerator{
		key: conf.GetSigningKey(),
	}
}

func (self *JWTGenerator) GetToken(issuer string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().UTC().Unix() + 1500,
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(self.key))
	return ss, err
}
