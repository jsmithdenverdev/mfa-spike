package api

import (
	"mfaspike/commands"
	"mfaspike/queries"

	"github.com/gorilla/mux"
)

type Api struct {
	Router   *mux.Router
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCode commands.CreateCodeHandler
	VerifyCode commands.VerifyCodeHandler
	CreateUser commands.CreateUserHandler
	DeleteUser commands.DeleteUserHandler
}

type Queries struct {
	ReadUser queries.ReadUserHandler
}
