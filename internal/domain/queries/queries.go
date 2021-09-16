package queries

import "mfaspike/internal/domain"

type userReader interface {
	Read(contact string) (domain.User, error)
}
