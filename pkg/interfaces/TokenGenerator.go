package interfaces

type SigningMethod interface {
	GetToken(string) (string, error)
}

// Public constructor
func NewTokenGenerator(signingMethod SigningMethod) *tokenGenerator {
	return &tokenGenerator{
		signingMethod: signingMethod,
	}
}

type tokenGenerator struct {
	signingMethod SigningMethod
}

func (self *tokenGenerator) GetToken(issuer string) (string, error) {
	token, err := self.signingMethod.GetToken(issuer)
	return token, err
}
