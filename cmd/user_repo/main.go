package main

import (
	"flag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"kkako_video/pkg/app"
	"kkako_video/pkg/db/mysqlx"
	"log"
)

func newApp(grpcServer *grpc.Server) (*app.App, error) {
	return app.NewApp(
		app.Name(appName),
		app.Port("9001"),
		app.GrpcServer(grpcServer),
	)
}

var conf = new(bool)
var appName = "user_repo"

func main() {
	flag.BoolVar(conf, "conf", false, "false is from file, true is from env")
	flag.Parse()
	if *conf == false {
		viper.AddConfigPath("configs")
		viper.SetConfigName(appName)
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln("read conf err:", err)
		}
	} else {
		viper.AutomaticEnv()
	}

	_, err := mysqlx.New(viper.GetViper())
	if err != nil {
		log.Fatalln("open mysql err:", err)
	}

	app, err := initApp(viper.GetViper())
	if err != nil {
		log.Fatalln("init app err:", err)
	}
	err = app.Start()
	if err != nil {
		log.Fatalln("app start err:", err)
	}
}
