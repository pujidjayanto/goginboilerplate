package user

import "errors"

var (
	ErrUserAlreadyExisted = errors.New("user already exist")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredential  = errors.New("invalid credential")
)
