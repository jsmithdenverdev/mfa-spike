package commands

// create code -----------------------

type CreateCode struct{}

type CreateCodeRequest struct {
	Contact string
}

type CreateCodeResponse struct {
	Contact string
	Code    string
}

func (cc CreateCode) Handle(request CreateCodeRequest) (CreateCodeResponse, error) {
	return CreateCodeResponse{}, nil
}

// verify code -----------------------

type VerifyCode struct{}

type VerifyCodeRequest struct {
	Contact string
	Code    string
}

type VerifyCodeResponse struct{}

func (vc VerifyCode) Handle(request VerifyCodeRequest) (VerifyCodeResponse, error) {
	return VerifyCodeResponse{}, nil
}

// expire code -----------------------

type ExpireCode struct{}

type ExpireCodeRequest struct{}

func (ec ExpireCode) Handle(request ExpireCodeRequest) error {
	return nil
}
