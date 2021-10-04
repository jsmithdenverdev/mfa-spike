package commands

import (
	"mfaspike"
	"time"
)

// create user -----------------------

type CreateUserHandler struct {
	writer userWriter
}

type CreateUserRequest struct {
	Contact  string
	Name     string
	Timezone string
}

type CreateUserResponse struct{}

func NewCreateUserHandler(writer userWriter) CreateUserHandler {
	return CreateUserHandler{
		writer,
	}
}

func (h *CreateUserHandler) Handle(request CreateUserRequest) (CreateUserResponse, error) {
	tz, err := time.LoadLocation(request.Timezone)

	if err != nil {
		// log the error but continue
	}

	err = h.writer.Write(&mfaspike.User{
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

type DeleteUserHandler struct {
	deleter userDeleter
}

type DeleteUserRequest struct {
	Contact string
}

type DeleteUserResponse struct {
}

func NewDeleteUserHandler(deleter userDeleter) DeleteUserHandler {
	return DeleteUserHandler{
		deleter,
	}
}

func (h *DeleteUserHandler) Handle(request DeleteUserRequest) (DeleteUserResponse, error) {
	err := h.deleter.Delete(request.Contact)

	if err != nil {
		return DeleteUserResponse{}, err
	}

	return DeleteUserResponse{}, nil
}
