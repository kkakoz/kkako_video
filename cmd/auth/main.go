package main

import (
	"context"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"kkako_video/internal/auth"
	"kkako_video/internal/auth/server"
	"kkako_video/internal/pkg/client"
	"kkako_video/pkg/app"
	"kkako_video/pkg/conf"
	"kkako_video/pkg/db/mysqlx"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func newApp(ctx context.Context, cancel context.CancelFunc, grpcServer *grpc.Server, httpServer http.Handler) (*app.App, error) {
	options := make([]app.Option, 0)
	options = append(options, app.GrpcServer(grpcServer), app.HttpServer(httpServer))
	return app.NewApp(
		ctx,
		cancel,
		options...,
	)
}

func main() {

	config := conf.ParseConf()
	ctx, cancel := context.WithCancel(context.TODO())

	_, err := mysqlx.New(config)
	if err != nil {
		log.Fatalln("open mysql err:", err)
	}
	var app = new(app.App)
	fx.New(
		auth.Provider,
		fx.Supply(config),
		fx.Provide(func() (context.Context, context.CancelFunc) {
			return ctx, cancel
		}),
		client.Provider,
		server.Provider,
		fx.Provide(newApp),
		fx.Populate(&app),
	)
	// 用于捕获退出信号
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	go func() {
		<-quit
		cancel()
	}()
	if err = app.Start(); err != nil {
		log.Fatal(err)
	}
}
