package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, WrapError(ErrInvalidSecretKeySize, "must be at least 32 bytes")
	}

	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token with the specified username and duration.
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString([]byte(maker.secretKey))
}

// VerifyToken checks if the token is valid or not.
func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		// Check the signing method to guarantee the trivial case.
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, WrapError(ErrInvalidToken, "unexpected signing method")
		}

		return []byte(maker.secretKey), nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := token.Claims.(*Payload)
	if !ok {
		return nil, WrapError(ErrInvalidToken, "invalid token claims")
	}

	return payload, nil
}
