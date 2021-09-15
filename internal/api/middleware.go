package api

import (
	"errors"
	"mfaspike/internal/domain"
	"mfaspike/internal/domain/commands"
	"net/http"
)

var (
	errNoPhone = errors.New("this endpoint requires an x-mfa-phone header")
	errNoCode  = errors.New("no mfa code")
)

func (a *Api) withMFA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mfaPhone := r.Header.Get("x-mfa-phone")
		mfaCode := r.Header.Get("x-mfa-code")

		// no mfa phone present on the request - invalid request
		if mfaPhone == "" {
			http.Error(w, errNoPhone.Error(), http.StatusUnauthorized)
			return
		}

		// no mfa code present on the request
		if mfaCode == "" {
			http.Error(w, errNoCode.Error(), http.StatusUnauthorized)
			w.WriteHeader(http.StatusUnauthorized)

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
			if errors.Is(err, domain.ErrExpiredCode) {
				http.Error(w, err.Error(), http.StatusUnauthorized)

				// re-initiate the mfa flow
				a.Commands.CreateCode.Handle(commands.CreateCodeRequest{
					Contact: mfaPhone,
				})

				return
			}

			if errors.Is(err, domain.ErrCodeMismatch) {
				http.Error(w, err.Error(), http.StatusUnauthorized)

				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
