package app_errors

import "errors"

var (
	EventNotFoundError      = errors.New("event not found")
	InvalidCredentialsError = errors.New("invalid credentials")
)
