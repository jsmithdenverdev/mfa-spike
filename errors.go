package mfaspike

import "errors"

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrExpiredCode  = errors.New("expired mfa code")
	ErrCodeMismatch = errors.New("mismatched mfa code")
	ErrNoCode       = errors.New("no mfa code")
)
