package commands

import (
	"mfaspike/internal/domain"
	"time"
)

// create user -----------------------

type CreateUser struct {
	Writer userWriter
}

type CreateUserRequest struct {
	Contact  string
	Name     string
	Timezone string
}

type CreateUserResponse struct{}

func (c *CreateUser) Handle(request CreateUserRequest) (CreateUserResponse, error) {
	tz, err := time.LoadLocation(request.Timezone)

	if err != nil {
		// log the error but continue
	}

	err = c.Writer.Write(&domain.User{
		Contact:  request.Contact,
		Name:     request.Name,
		Timezone: *tz,
	})

	if err != nil {
		return CreateUserResponse{}, nil
	}

	return CreateUserResponse{}, nil
}

// delete user -----------------------

type DeleteUserRequest struct {
	Contact string
}

type DeleteUserResponse struct{}

type DeleteUser struct {
	Deleter userDeleter
}

func (c *DeleteUser) Handle(request DeleteUserRequest) (DeleteUserResponse, error) {
	err := c.Deleter.Delete(request.Contact)

	if err != nil {
		return DeleteUserResponse{}, err
	}

	return DeleteUserResponse{}, nil
}
