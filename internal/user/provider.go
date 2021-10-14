package user

import (
	"go.uber.org/fx"
	"kkako_video/internal/pkg/jwtx"
	"kkako_video/internal/user/handler"
	"kkako_video/internal/user/logic"
	"kkako_video/internal/user/repo"
	"kkako_video/internal/user/server"
)

var Provider = fx.Provide(handler.NewUserHandler, logic.NewUserLogic, repo.NewUserRepo, server.NewGrpcServer, jwtx.NewJwtTokenVerifier)

