package main

import (
	"context"
	"flag"
	"github.com/pkg/errors"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"kkako_video/internal/auth"
	"kkako_video/internal/pkg/client"
	"kkako_video/pkg/app"
	"kkako_video/pkg/conf"
	"kkako_video/pkg/db/mysqlx"
	"log"
	"os"
)

func newApp(ctx context.Context, grpcServer *grpc.Server) (*app.App, error) {
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
	options = append(options, app.Port("9001"), app.GrpcServer(grpcServer))
	return app.NewApp(
		ctx,
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
		fx.Provide(func() context.Context {
			return ctx
		}),
		client.Provider,
		fx.Provide(newApp),
		fx.Populate(&app),
	)
	defer cancel()
	if err := app.Start(); err != nil {  // 手动调用Start
		log.Fatal(err)
	}
}
