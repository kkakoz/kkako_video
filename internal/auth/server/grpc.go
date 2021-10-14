package server

import (
	"go.uber.org/fx"
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/auth/handler"
)

func NewGrpcServer(handler *handler.AuthHandler) *grpc.Server {
	server := grpc.NewServer()
	v1.RegisterAuthServiceServer(server, handler)
	return server
}


var Provider = fx.Provide(NewGrpcServer)