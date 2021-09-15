package domain

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrExpiredCode  = errors.New("expired mfa code")
	ErrCodeMismatch = errors.New("mismatched mfa code")
)
