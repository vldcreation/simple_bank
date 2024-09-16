package token

import "errors"

var (
	ErrInvalidSecretKeySize = errors.New("invalid secret key size")
	ErrExpiredToken         = errors.New("token is expired")
	ErrInvalidToken         = errors.New("invalid token")
)

func WrapError(err error, msg string) error {
	return errors.New(msg + ": " + err.Error())
}
