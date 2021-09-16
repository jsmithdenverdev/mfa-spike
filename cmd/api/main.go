package main

import (
	"mfaspike/internal/api"
	"mfaspike/internal/domain/commands"
	"mfaspike/internal/domain/queries"
	"mfaspike/internal/storage"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	userDb, err := gorm.Open(sqlite.Open("user.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	mfaDb, err := gorm.Open(sqlite.Open("mfa.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	userStore := storage.UserStore{
		Client: userDb,
	}

	mfaStore := storage.MfaStore{
		Client: mfaDb,
	}

	api := api.Api{
		Commands: api.Commands{
			CreateUser: commands.CreateUser{
				Writer: &userStore,
			},
			DeleteUser: commands.DeleteUser{
				Deleter: &userStore,
			},
			CreateCode: commands.CreateCode{
				Writer: &mfaStore,
			},
			VerifyCode: commands.VerifyCode{
				Reader: &mfaStore,
			},
		},
		Queries: api.Queries{
			ReadUser: queries.ReadUser{
				Reader: &userStore,
			},
		},
	}

	api.Router = mux.NewRouter()
	api.Routes()

	http.ListenAndServe("localhost:5000", api.Router)
}
