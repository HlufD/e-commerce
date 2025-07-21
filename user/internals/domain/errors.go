package domain

import "errors"

var (
	// User errors
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExists        = errors.New("email already exists")
	ErrUsernameExists     = errors.New("username already exists")

	// Token errors
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("token has expired")

	// Validation errors
	ErrInvalidEmail = errors.New("invalid email format")
	ErrWeakPassword = errors.New("password does not meet requirements")
)
