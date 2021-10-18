package main

import (
	"context"
	"flag"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
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
	var debug bool
	flag.BoolVar(&debug, "debug", false, "false is from file, true is from env")
	flag.Parse()
	if !debug {
		name, b := os.LookupEnv("MY_POD_NAME")
		if !b {
			return nil, errors.New("get name err")
		}
		options = append(options, app.Name(name))
		ip, b := os.LookupEnv("MY_POD_IP")
		if !b {
			return nil, errors.New("get ip err")
		}
		options = append(options, app.IP(ip))
	}
	options = append(options, app.Port(viper.GetString("app.port")), app.GrpcServer(grpcServer), app.HttpServer(httpServer))
	return app.NewApp(
		ctx,
		cancel,
		options...,
	)
}

func main() {

	viper := conf.ParseConf()
	ctx, cancel := context.WithCancel(context.TODO())

	_, err := mysqlx.New(viper)
	if err != nil {
		log.Fatalln("open mysql err:", err)
	}
	var app = new(app.App)
	fx.New(
		auth.Provider,
		fx.Supply(viper),
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
