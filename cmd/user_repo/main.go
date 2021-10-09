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
	options := make([]app.Option, 0)
	if !*debug {
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
		options...
	)
}

var debug = new(bool)

func main() {
	flag.BoolVar(debug, "debug", false, "false is from file, true is from env")
	flag.Parse()
	if *debug {
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
