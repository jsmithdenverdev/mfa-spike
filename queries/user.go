package queries

import "mfaspike"

// read user -----------------------

type ReadUserHandler struct {
	reader userReader
}

type ReadUserRequest struct {
	Contact string
}

type ReadUserResponse struct {
	User mfaspike.User
}

func NewReadUserHandler(reader userReader) ReadUserHandler {
	return ReadUserHandler{
		reader,
	}
}

func (h *ReadUserHandler) Handle(request ReadUserRequest) (ReadUserResponse, error) {
	user, err := h.reader.Read(request.Contact)

	if err != nil {
		return ReadUserResponse{}, err
	}

	return ReadUserResponse{
		User: user,
	}, nil
}
