package user

import (
	"go.uber.org/fx"
	"kkako_video/internal/user/handler"
	"kkako_video/internal/user/logic"
	"kkako_video/internal/user/repo"
)

var Provider = fx.Provide(handler.NewUserHandler, logic.NewUserLogic, repo.NewUserRepo, repo.NewAuthRepo)

