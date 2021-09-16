package commands

import (
	"math/rand"
	"mfaspike/internal/domain"
	"strconv"
	"time"
)

// create code -----------------------

type CreateCode struct {
	Writer mfaCodeWriter
}

type CreateCodeRequest struct {
	Contact string
}

type CreateCodeResponse struct {
	Contact string
	Code    string
}

func (cc CreateCode) Handle(request CreateCodeRequest) (CreateCodeResponse, error) {
	var (
		seed   = rand.New(rand.NewSource(time.Now().Unix()))
		random = strconv.Itoa(seed.Int())
		code   = random[0:6]
	)

	err := cc.Writer.Write(&domain.MfaCode{
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

type VerifyCode struct {
	Reader mfaCodeReader
}

type VerifyCodeRequest struct {
	Contact string
	Code    string
}

type VerifyCodeResponse struct{}

func (vc VerifyCode) Handle(request VerifyCodeRequest) (VerifyCodeResponse, error) {
	code, err := vc.Reader.Read(request.Contact)

	// TODO: this error might be that there are no records
	if err != nil {
		return VerifyCodeResponse{}, nil
	}

	if code.Code != request.Code {
		return VerifyCodeResponse{}, domain.ErrCodeMismatch
	}

	return VerifyCodeResponse{}, nil
}

// expire code -----------------------

type ExpireCode struct {
	Deleter mfaCodeDeleter
}

type ExpireCodeRequest struct{}

func (ec ExpireCode) Handle(request ExpireCodeRequest) error {
	return nil
}
