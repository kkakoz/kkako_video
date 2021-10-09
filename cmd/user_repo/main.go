package main

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"kkako_video/pkg/app"
	"kkako_video/pkg/db/mysqlx"
	"log"
	"os"
)

func newApp(grpcServer *grpc.Server) (*app.App, error) {
	name, b := os.LookupEnv("MY_POD_NAME")
	if !b {
		return nil, errors.New("get name err")
	}
	ip, b := os.LookupEnv("MY_POD_IP")
	if !b {
		return nil, errors.New("get ip err")
	}
	return app.NewApp(
		app.Name(name),
		app.IP(ip),
		app.Port("9001"),
		app.GrpcServer(grpcServer),
	)
}

var debug = new(bool)

func main() {
	flag.BoolVar(debug, "debug", true, "false is from file, true is from env")
	flag.Parse()
	if *debug == true {
		viper.AddConfigPath("configs")
		viper.SetConfigName("test")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln("read conf err:", err)
		}
	} else {
		viper.AddConfigPath("configs")
		viper.SetConfigName("conf")
		viper.SetConfigType("yaml")
		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalln("read conf err:", err)
		}
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
