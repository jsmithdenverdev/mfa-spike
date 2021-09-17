package main

import (
	"log"
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

	userStore := storage.NewUserStore(userDb)
	mfaStore := storage.NewMfaStore(mfaDb)

	api := api.Api{
		Commands: api.Commands{
			CreateUser: commands.NewCreateUserHandler(&userStore),
			DeleteUser: commands.NewDeleteUserHandler(&userStore),
			CreateCode: commands.NewCreateCodeHandler(&mfaStore),
			VerifyCode: commands.NewVerifyCodeHandler(&mfaStore),
		},
		Queries: api.Queries{
			ReadUser: queries.NewReadUserHandler(&userStore),
		},
	}

	api.Router = mux.NewRouter()

	api.ConfigureRoutes()

	log.Fatal(http.ListenAndServe("localhost:5000", api.Router))
}
