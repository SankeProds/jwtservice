package implementation

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTConfiguration interface {
	GetSigningKey() string
}

type SigningKeyGetter interface {
	GetSigningKey() []byte
}

// Public constructor
func NewJWTGenerator(keyGetter SigningKeyGetter) *jwtGenerator {
	return &jwtGenerator{
		keyGetter: keyGetter,
	}
}

type jwtGenerator struct {
	keyGetter SigningKeyGetter
}

func (jwtGen *jwtGenerator) GetToken(issuer string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(jwtGen.keyGetter.GetSigningKey())
	return ss, err
}

func NewSigningKeyGetter(conf JWTConfiguration) SigningKeyGetter {
	return &signingKeyGetter{
		// This is not a good idea, better use a vault services
		// Its more secure
		// also, if one of your signing keys changes all your services
		// you just need to update the vault
		base: []byte(conf.GetSigningKey()),
	}
}

type signingKeyGetter struct {
	base []byte
}

func (keyGetter *signingKeyGetter) GetSigningKey() []byte {
	// signing key changes each hour
	// is this a good idea? IDK
	dinamic := time.Now().Unix() / 60 / 60 // hours
	return []byte(fmt.Sprintf("%s%+v", keyGetter.base, dinamic))
}
