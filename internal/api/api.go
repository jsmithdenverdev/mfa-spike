package api

import "mfaspike/internal/domain/commands"

type Api struct {
	Commands Commands
}

type Commands struct {
	CreateCode commands.CreateCode
	VerifyCode commands.VerifyCode
}
