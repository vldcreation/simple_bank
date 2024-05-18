package util

import "strings"

// ErrKeyVal is a simple error type with a key and a value
type ErrKeyVal struct {
	k string
	v string
}

func UnwrapError(err string) ErrKeyVal {
	if err == "" {
		return ErrKeyVal{}
	}

	// split the error message into key and value
	errParts := strings.Split(err, ":")
	if len(errParts) != 2 {
		return ErrKeyVal{}
	}

	return ErrKeyVal{
		k: strings.TrimSpace(errParts[0]),
		v: strings.TrimSpace(errParts[1]),
	}
}

func WrapError(k, v string) error {
	return ErrKeyVal{
		k: k,
		v: v,
	}
}

func (e ErrKeyVal) Error() string {
	return e.k + ": " + e.v
}

func (e ErrKeyVal) Is(err error) bool {
	other, ok := err.(ErrKeyVal)
	if !ok {
		return false
	}
	return e.k == other.k && e.v == other.v
}

func (e ErrKeyVal) Value() string {
	return e.v
}

func (e ErrKeyVal) Key() string {
	return e.k
}
