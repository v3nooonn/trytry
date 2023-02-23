package jwtx

import (
	"crypto/ecdsa"
	"time"

	e "github.com/v3nooonn/trytry/pkg/errorx"

	"github.com/dgrijalva/jwt-go"
)

type (
	// AssignerOpt functional options of assigner generation
	AssignerOpt func(assigner *Assigner)

	OptFunc func() []AssignerOpt
)

type (
	Assigner struct {
		jwt.StandardClaims

		PrivatePEM string
		Expiration time.Duration

		Temporary *struct{}
		Access    *struct{}
		Refresh   *struct{}
	}

	Claims struct {
		jwt.StandardClaims
		// Temporary
		// Access
		// Refresh
	}
)

func NewAssigner(std jwt.StandardClaims, pem string, exp time.Duration, optFunc OptFunc) (*Assigner, error) {
	assigner := &Assigner{
		StandardClaims: std,
		PrivatePEM:     pem,
		Expiration:     exp,
	}

	if optFunc == nil {
		return assigner, nil
	}

	for _, opt := range optFunc() {
		opt(assigner)
	}

	return assigner, nil
}

// Assign jwt assignment
func (a *Assigner) Assign() (string, error) {
	rsaKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(a.PrivatePEM))
	if err != nil {
		return "", e.Internal("PEM key validation error")
	}

	now := time.Now()

	a.StandardClaims.IssuedAt = now.Unix()
	a.StandardClaims.ExpiresAt = now.Add(a.Expiration).Unix()
	claim := Claims{StandardClaims: a.StandardClaims}

	token := jwt.NewWithClaims(jwt.GetSigningMethod(jwt.SigningMethodES256.Name), claim)

	return assign(*token, rsaKey)
}

func assign(token jwt.Token, key *ecdsa.PrivateKey) (string, error) {
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", e.Internal("signed string generation error")
	}

	return tokenString, nil
}
