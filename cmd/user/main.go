package main

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"kkako_video/internal/pkg/client"
	"kkako_video/internal/user"
	"kkako_video/pkg/app"
	"kkako_video/pkg/conf"
	"kkako_video/pkg/db/mysqlx"
	"log"
	"os"
)

func newApp(ctx context.Context, cancel context.CancelFunc, grpcServer *grpc.Server) (*app.App, error) {
	options := make([]app.Option, 0)
	debug := viper.GetViper().GetBool("app.debug")
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
	options = append(options, app.Port(viper.GetString("app.port")), app.GrpcServer(grpcServer))
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
		user.Provider,
		fx.Supply(config),
		fx.Provide(func() (context.Context, context.CancelFunc) {
			return ctx, cancel
		}),
		client.Provider,
		fx.Provide(newApp),
		fx.Populate(&app),
	)
	defer cancel()
	if err = app.Start(); err != nil {  // 手动调用Start
		log.Fatal(err)
	}
}
