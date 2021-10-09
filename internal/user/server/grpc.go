package server

import (
	"github.com/google/wire"
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user/handler"
)

func NewGrpcServer(handler *handler.UserHandler) *grpc.Server {
	server := grpc.NewServer()
	v1.RegisterLoginServiceServer(server, handler)
	v1.RegisterUserServiceServer(server, handler)
	return server
}


var Provider = wire.NewSet(NewGrpcServer)