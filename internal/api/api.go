package api

import (
	"mfaspike/internal/domain/commands"
	"mfaspike/internal/domain/queries"

	"github.com/gorilla/mux"
)

type Api struct {
	Router   *mux.Router
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCode commands.CreateCode
	VerifyCode commands.VerifyCode
	CreateUser commands.CreateUser
	DeleteUser commands.DeleteUser
}

type Queries struct {
	ReadUser queries.ReadUser
}
