package server

import (
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
)

func NewGrpcServer(handler v1.UserRepoServiceServer) *grpc.Server {
	server := grpc.NewServer()
	v1.RegisterUserRepoServiceServer(server, handler)
	return server
}
