package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user/handler"
	"net/http"
)

func NewGrpcServer(handler *handler.UserHandler) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	v1.RegisterUserServiceServer(server, handler)
	return server
}

func NewHttpServer() (http.Handler, error) {
	port := viper.GetString("app.port")
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := v1.NewUserServiceClient(conn)
	gwmux := runtime.NewServeMux()
	err = v1.RegisterUserServiceHandlerClient(context.TODO(), gwmux, client)
	if err != nil {
		return gwmux, err
	}
	return gwmux, nil
}

var Provider = fx.Provide(NewGrpcServer, NewHttpServer)