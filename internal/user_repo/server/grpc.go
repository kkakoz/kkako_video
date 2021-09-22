package server

import (
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
	handler2 "kkako_video/internal/user_repo/handler"
)

func NewGrpcServer(handler handler2.UserRepoHandler) *grpc.Server {
	server := grpc.NewServer()
	v1.RegisterUserRepoServiceServer(server, handler)
	return server
}
