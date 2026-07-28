package jwt

import "errors"

var (
	ErrInvalidToken            = errors.New("invalid token")
	ErrExpiredToken            = errors.New("token has expired")
	ErrInvalidSubject          = errors.New("invalid token subject")
	ErrInvalidRole             = errors.New("invalid role")
	ErrInsufficientPermissions = errors.New("insufficient permissions")
)
