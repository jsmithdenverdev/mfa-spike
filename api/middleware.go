package api

import (
	"errors"
	"mfaspike"
	"mfaspike/commands"
	"net/http"
)

var (
	errNoPhone = errors.New("this endpoint requires an X-MFA-Phone header")
	errNoCode  = errors.New("this endpoint requires an X-MFA-Code header")
)

func (a *Api) withMFA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mfaPhone := r.Header.Get("X-MFA-Phone")
		mfaCode := r.Header.Get("X-MFA-Code")

		// no mfa phone present on the request - invalid request
		if mfaPhone == "" {
			http.Error(w, errNoPhone.Error(), http.StatusUnauthorized)

			return
		}

		// no mfa code present on the request
		if mfaCode == "" {
			http.Error(w, errNoCode.Error(), http.StatusUnauthorized)

			// initiate the mfa flow
			a.Commands.CreateCode.Handle(commands.CreateCodeRequest{
				Contact: mfaPhone,
			})

			return
		}

		_, err := a.Commands.VerifyCode.Handle(commands.VerifyCodeRequest{
			Contact: mfaPhone,
			Code:    mfaCode,
		})

		if err != nil {
			if errors.Is(err, mfaspike.ErrExpiredCode) {
				http.Error(w, err.Error(), http.StatusUnauthorized)

				// re-initiate the mfa flow
				a.Commands.CreateCode.Handle(commands.CreateCodeRequest{
					Contact: mfaPhone,
				})

				return
			}

			if errors.Is(err, mfaspike.ErrCodeMismatch) {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			http.Error(w, ErrInternalServer.Error(), http.StatusInternalServerError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
