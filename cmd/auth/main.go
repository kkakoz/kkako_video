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

func invokeApp(ctx context.Context, grpcServer *grpc.Server) error {
	options := make([]app.Option, 0)
	var debug bool
	flag.BoolVar(&debug, "debug", false, "false is from file, true is from env")
	flag.Parse()
	if !debug {
		name, b := os.LookupEnv("MY_POD_NAME")
		if !b {
			return errors.New("get name err")
		}
		options = append(options, app.Name(name))
		ip, b := os.LookupEnv("MY_POD_IP")
		if !b {
			return errors.New("get ip err")
		}
		options = append(options, app.IP(ip))
	}
	options = append(options, app.Port("9001"), app.GrpcServer(grpcServer))
	app, err := app.NewApp(
		ctx,
		options...,
	)
	if err != nil {
		return err
	}
	err = app.Start()
	return err
}

func main() {

	viper := conf.ParseConf()
	ctx, cancel := context.WithCancel(context.TODO())

	_, err := mysqlx.New(viper)
	if err != nil {
		log.Fatalln("open mysql err:", err)
	}
	app := fx.New(
		auth.Provider,
		fx.Supply(viper),
		fx.Provide(func() context.Context {
			return ctx
		}),
		client.Provider,
		fx.Invoke(invokeApp),
	)
	defer cancel()
	if err := app.Start(ctx); err != nil {  // 手动调用Start
		log.Fatal(err)
	}
}
