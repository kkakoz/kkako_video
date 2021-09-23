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
	)
}

var conf *bool
var appName = "user_repo"

func main() {
	flag.BoolVar(conf, "conf", false, "false is from file, true is from env")
	flag.Parse()
	if *conf == false {
		viper.AddConfigPath("../config")
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

}
