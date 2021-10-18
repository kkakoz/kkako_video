package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/auth/handler"
	"net/http"
)

func NewGrpcServer(handler *handler.AuthHandler) *grpc.Server {
	server := grpc.NewServer()
	v1.RegisterAuthServiceServer(server, handler)
	return server
}

func NewHttpServer(handler *handler.AuthHandler) (http.Handler, error) {
	gwmux := runtime.NewServeMux(
		//runtime.WithErrorHandler(NewHandleErr(logger)),
		)
	err := v1.RegisterAuthServiceHandlerServer(context.TODO(), gwmux, handler)
	if err != nil {
		return gwmux, err
	}
	return gwmux, nil
}


var Provider = fx.Provide(NewGrpcServer, NewHttpServer)