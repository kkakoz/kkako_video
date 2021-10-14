package auth

import (
	"go.uber.org/fx"
	"kkako_video/internal/auth/handler"
	"kkako_video/internal/auth/logic"
	"kkako_video/internal/auth/repo"
	"kkako_video/internal/auth/server"
	"kkako_video/internal/pkg/jwtx"
)

var Provider = fx.Provide(handler.NewAuthHandler, logic.NewAuthLogic, repo.NewAuthRepo, server.NewGrpcServer, jwtx.NewJwtTokenGen)
