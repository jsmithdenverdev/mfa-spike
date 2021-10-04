package main

import (
	"log"
	"mfaspike/api"
	"mfaspike/cache"
	"mfaspike/commands"
	"mfaspike/queries"
	"mfaspike/storage"
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

	userStore, err := storage.NewUserStore(userDb)
	if err != nil {
		panic(err)
	}

	mfaCache := cache.NewCodeCache()

	api := api.Api{
		Commands: api.Commands{
			CreateUser: commands.NewCreateUserHandler(&userStore),
			DeleteUser: commands.NewDeleteUserHandler(&userStore),
			CreateCode: commands.NewCreateCodeHandler(&mfaCache),
			VerifyCode: commands.NewVerifyCodeHandler(&mfaCache),
		},
		Queries: api.Queries{
			ReadUser: queries.NewReadUserHandler(&userStore),
		},
	}

	api.Router = mux.NewRouter()

	api.ConfigureRoutes()

	log.Fatal(http.ListenAndServe("localhost:5000", api.Router))
}
