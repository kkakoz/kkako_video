package conf

import (
	"github.com/spf13/viper"
	"log"
)

var debug bool

func Debug() bool {
	return debug
}

func ParseConf() *viper.Viper {
	viper.AddConfigPath("configs")
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read conf errerr:", err)
	}
	return viper.GetViper()
}
