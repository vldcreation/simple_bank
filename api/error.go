package api

import "net/http"

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ErrorOpt func(*ApiError)

func WithMessage(msg string) ErrorOpt {
	return func(e *ApiError) {
		e.Message = msg
	}
}

func NewApiError(code int, opt ...ErrorOpt) *ApiError {
	e := &ApiError{
		Code:    code,
		Message: http.StatusText(code),
	}

	for _, o := range opt {
		o(e)
	}
	return e
}
func (e *ApiError) WithMessage(msg string) *ApiError {
	e.Message = msg
	return e
}

func (e *ApiError) Error() string {
	return e.Message
}
