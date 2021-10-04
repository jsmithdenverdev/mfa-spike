package commands

import (
	"log"
	"math/rand"
	"mfaspike"
	"strconv"
	"time"
)

// create code -----------------------

type CreateCodeHandler struct {
	writer mfaCodeWriter
}

type CreateCodeRequest struct {
	Contact string
}

type CreateCodeResponse struct {
	Contact string
	Code    string
}

func NewCreateCodeHandler(writer mfaCodeWriter) CreateCodeHandler {
	return CreateCodeHandler{
		writer,
	}
}

func (h CreateCodeHandler) Handle(request CreateCodeRequest) (CreateCodeResponse, error) {
	var (
		seed   = rand.New(rand.NewSource(time.Now().Unix()))
		random = strconv.Itoa(seed.Int())
		code   = random[0:6]
	)

	log.Printf("mfa code: %s", code)

	err := h.writer.Write(&mfaspike.Code{
		Contact: request.Contact,
		Code:    code,
	})

	if err != nil {

		return CreateCodeResponse{}, err
	}

	return CreateCodeResponse{
		Contact: request.Contact,
		Code:    code,
	}, nil
}

// verify code -----------------------

type VerifyCodeHandler struct {
	reader mfaCodeReader
}

type VerifyCodeRequest struct {
	Contact string
	Code    string
}

type VerifyCodeResponse struct{}

func NewVerifyCodeHandler(reader mfaCodeReader) VerifyCodeHandler {
	return VerifyCodeHandler{
		reader,
	}
}

func (h VerifyCodeHandler) Handle(request VerifyCodeRequest) (VerifyCodeResponse, error) {
	code, err := h.reader.Read(request.Contact)

	// TODO: this error might be that there are no records
	if err != nil {
		return VerifyCodeResponse{}, err
	}

	if code.Code != request.Code {
		return VerifyCodeResponse{}, mfaspike.ErrCodeMismatch
	}

	return VerifyCodeResponse{}, nil
}

// expire code -----------------------

type ExpireCodeHandler struct {
	deleter mfaCodeDeleter
}

type ExpireCodeRequest struct {
}

func NewExpireCodeHandler(deleter mfaCodeDeleter) ExpireCodeHandler {
	return ExpireCodeHandler{
		deleter,
	}
}

func (h ExpireCodeHandler) Handle(request ExpireCodeRequest) error {
	return nil
}
