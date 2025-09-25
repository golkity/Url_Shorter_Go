package custom_errors

import "errors"

var (
	ErrUserExists    = errors.New("user with this email already exists")
	ErrInvalidCreds  = errors.New("invalid credentials")
	ErrPasswordShort = errors.New("password is too short")
)
