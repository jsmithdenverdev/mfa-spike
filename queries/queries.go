package queries

import "mfaspike"

type userReader interface {
	Read(contact string) (mfaspike.User, error)
}
