package exception

import (
	"errors"
)

var (
	ErrPermissionDenied    = errors.New("permission denied")
	ErrRecordNotFound      = errors.New("record not found")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)

type ErrorSendToResponse struct {
	Err string
}

func (e *ErrorSendToResponse) Error() string {
	return e.Err
}
