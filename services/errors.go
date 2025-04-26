package services

import "errors"

var (
	ErrEmailExists    = errors.New("email already registered")
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidToken   = errors.New("invalid token")
	ErrInvalidSession = errors.New("invalid session")
	ErrSessionExpired = errors.New("session expired")
	ErrWrongPassword  = errors.New("wrong password")
)
