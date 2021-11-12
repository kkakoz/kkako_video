package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	v1 "kkako_video/api/video/v1"
	"kkako_video/internal/pkg/jwtx"
	"kkako_video/internal/video/handler"
	"net/http"
)

func NewGrpcServer(handler *handler.VideoHandler, verifier *jwtx.JwtTokenVerifier) *grpc.Server {
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(),
	)
	v1.RegisterVideoServiceServer(server, handler)
	return server
}

func NewHttpServer(handler *handler.VideoHandler) (http.Handler, error) {
	gwmux := runtime.NewServeMux()
	err := v1.RegisterVideoServiceHandlerServer(context.TODO(), gwmux, handler)
	if err != nil {
		return gwmux, err
	}
	return gwmux, nil
}



var Provider = fx.Provide(NewGrpcServer, NewHttpServer)