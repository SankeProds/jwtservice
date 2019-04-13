package implementation

import "github.com/dgrijalva/jwt-go"

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
		ExpiresAt: 15000,
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(self.key)
	return ss, err
}
