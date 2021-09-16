package commands

import "mfaspike/internal/domain"

type mfaCodeWriter interface {
	Write(code *domain.MfaCode) error
}

type mfaCodeReader interface {
	Read(contact string) (domain.MfaCode, error)
}

type mfaCodeDeleter interface {
	Delete(contact string) error
}

type userWriter interface {
	Write(code *domain.User) error
}

type userDeleter interface {
	Delete(contact string) error
}
