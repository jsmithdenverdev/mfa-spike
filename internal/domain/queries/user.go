package queries

import "mfaspike/internal/domain"

// read user -----------------------

type ReadUser struct {
	Reader userReader
}

type ReadUserRequest struct {
	Contact string
}

type ReadUserResponse struct {
	User domain.User
}

func (q *ReadUser) Handle(request ReadUserRequest) (ReadUserResponse, error) {
	user, err := q.Reader.Read(request.Contact)

	if err != nil {
		return ReadUserResponse{}, err
	}

	return ReadUserResponse{
		User: user,
	}, nil
}
