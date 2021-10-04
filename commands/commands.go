package commands

import "mfaspike"

type mfaCodeWriter interface {
	Write(code *mfaspike.Code) error
}

type mfaCodeReader interface {
	Read(contact string) (mfaspike.Code, error)
}

type mfaCodeDeleter interface {
	Delete(contact string) error
}

type userWriter interface {
	Write(code *mfaspike.User) error
}

type userDeleter interface {
	Delete(contact string) error
}
